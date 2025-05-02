package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterWiresharkTool(s *server.MCPServer) {
	// Add a Wireshark tool
	wiresharkTool := mcp.NewTool("capture_packets",
		mcp.WithDescription("Capture live traffic and provide raw packet data as JSON for LLM analysis"),
		mcp.WithString("interface",
			mcp.DefaultString("en0"),
			mcp.Description("Network interface to capture from (e.g., eth0, en0)"),
		),
		mcp.WithNumber("duration",
			mcp.DefaultNumber(10),
			mcp.Description("Capture duration in seconds"),
		),
	)

	s.AddTool(wiresharkTool, capturePackets)
}

func capturePackets(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	interfaceName := request.Params.Arguments["interface"].(string)
	duration := int(request.Params.Arguments["duration"].(float64))

	tsharkPath, err := findTshark()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error: %s", err.Error())), nil
	}

	tempPcap := "temp_capture.pcap"
	fmt.Printf("Capturing packets on %s for %d seconds\n", interfaceName, duration)

	// Run tshark to capture packets
	captureCmd := exec.Command(tsharkPath, "-i", interfaceName, "-w", tempPcap, "-a", fmt.Sprintf("duration:%d", duration))
	if err := captureCmd.Run(); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error capturing packets: %s", err.Error())), nil
	}

	// Read captured packets and convert to JSON
	readCmd := exec.Command(tsharkPath, "-r", tempPcap, "-T", "json",
		"-e", "frame.number", "-e", "ip.src", "-e", "ip.dst",
		"-e", "tcp.srcport", "-e", "tcp.dstport", "-e", "tcp.flags",
		"-e", "frame.time", "-e", "http.request.method", "-e", "http.response.code")
	readCmd.Env = append(os.Environ(), "PATH=/usr/bin:/usr/local/bin:/opt/homebrew/bin")
	output, err := readCmd.Output()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error reading packets: %s", err.Error())), nil
	}

	// Parse JSON and trim if necessary
	var packets []map[string]interface{}
	if err := json.Unmarshal(output, &packets); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error parsing JSON: %s", err.Error())), nil
	}

	jsonString, err := json.Marshal(packets)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error serializing JSON: %s", err.Error())), nil
	}

	// Clean up temporary file
	if err := os.Remove(tempPcap); err != nil {
		fmt.Printf("Failed to delete %s: %s\n", tempPcap, err.Error())
	}

	return mcp.NewToolResultText(fmt.Sprintf("Captured packet data (JSON for LLM analysis):\n%s", string(jsonString))), nil
}

func findTshark() (string, error) {
	tsharkPath, err := exec.LookPath("tshark")
	if err == nil {
		fmt.Printf("Found tshark at: %s\n", tsharkPath)
		return tsharkPath, nil
	}

	fmt.Printf("Failed to find tshark using LookPath: %s\n", err.Error())
	fallbacks := []string{
		"/usr/bin/tshark",
		"/usr/local/bin/tshark",
		"/opt/homebrew/bin/tshark",
		"/Applications/Wireshark.app/Contents/MacOS/tshark",
	}

	for _, path := range fallbacks {
		cmd := exec.Command(path, "-v")
		if err := cmd.Run(); err == nil {
			fmt.Printf("Found tshark at fallback: %s\n", path)
			return path, nil
		} else {
			fmt.Printf("Fallback %s failed: %s\n", path, err.Error())
		}
	}

	return "", errors.New("tshark not found. Please install Wireshark (https://www.wireshark.org/download.html) and ensure tshark is in your PATH")
}
