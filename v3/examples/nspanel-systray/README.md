# NSPanel + System Tray Example

This example demonstrates the new NSPanel integration for Wails v3 on macOS, combined with system tray functionality.

## Features

- **NSPanel Integration**: Converts a regular window to a macOS NSPanel with native panel behaviors
- **System Tray**: System tray icon with context menu for easy panel management
- **Auto-conversion**: Automatically converts to NSPanel on macOS when the window loads
- **Panel Behaviors**: 
  - Floats above other windows
  - Hides when it becomes inactive
  - Becomes key only when needed
  - Proper resource management when closed
- **Cross-platform**: Gracefully handles non-macOS platforms (NSPanel features disabled)

## NSPanel Benefits

NSPanel provides several advantages over regular windows on macOS:

1. **Better UX**: Panels are designed for utility windows and have better focus management
2. **System Integration**: Properly integrates with macOS window management
3. **Performance**: More efficient memory and resource usage
4. **Behaviors**: Native panel behaviors like auto-hiding and floating

## How to Run

1. Navigate to the example directory:
   ```bash
   cd v3/examples/nspanel-systray
   ```

2. Run the example:
   ```bash
   go run .
   ```

3. On macOS:
   - The app will start with a system tray icon
   - Click the tray icon to show the panel
   - The window will automatically convert to an NSPanel
   - The panel will float above other windows and hide when inactive

4. On other platforms:
   - The app will work normally but without NSPanel features
   - You'll see a message indicating NSPanel is macOS-only

## Usage

### System Tray

- **Left click**: Toggle panel visibility
- **Right click**: Show context menu with options:
  - Show Panel
  - Toggle Panel
  - Quit

### Panel Controls

The panel interface provides several controls:

- **Convert to NSPanel**: Manually trigger panel conversion (auto-disabled after conversion)
- **Show Window**: Make the panel visible
- **Hide Window**: Hide the panel
- **Toggle Panel**: Toggle visibility
- **Refresh Status**: Update the status information

### Status Information

The panel displays:
- Current platform (OS/Architecture)
- Whether the window is currently a panel
- Go runtime version

## NSPanel API

The NSPanel integration adds the following methods to `WebviewWindow`:

### Conversion
```go
// Convert window to NSPanel (macOS only)
err := window.ConvertToPanel()

// Check if window is a panel
isPanel := window.IsPanel()
```

### Panel Configuration
```go
// Set panel to float above other windows
err := window.SetPanelFloating(true)

// Set panel to hide when it becomes inactive
err := window.SetPanelHidesOnDeactivate(true)

// Set panel to become key only when needed
err := window.SetPanelBecomesKeyOnlyIfNeeded(true)

// Set whether panel works when modal windows are present
err := window.SetPanelWorksWhenModal(false)

// Set whether panel is released when closed
err := window.SetPanelReleasedWhenClosed(true)
```

All NSPanel methods return appropriate errors on non-macOS platforms.

## Implementation Details

### Backend (Go)

The `AppService` provides methods to:
- Convert windows to panels
- Check panel status
- Control window visibility
- Get system information

### Frontend (HTML/JS)

The frontend provides a modern, responsive interface using:
- Native system fonts
- Gradient backgrounds with glassmorphism effects
- Real-time status updates
- Responsive design
- Error handling and notifications

### Platform Detection

The example automatically detects the platform and:
- Enables NSPanel features on macOS
- Disables NSPanel features on other platforms
- Shows appropriate notifications and warnings

## System Requirements

- **macOS**: Full NSPanel functionality
- **Windows/Linux**: System tray and basic window functionality (NSPanel features disabled)
- **Go**: 1.21 or later
- **Wails v3**: Latest version with NSPanel support

## Troubleshooting

### Common Issues

1. **Panel not converting**: 
   - Ensure you're running on macOS
   - Check console logs for errors
   - Try manual conversion using the button

2. **System tray not appearing**:
   - Ensure system tray is enabled in your OS
   - Check if the app has necessary permissions

3. **Panel not hiding**:
   - This is expected behavior - panels may stay visible based on focus rules
   - Use the hide button or system tray menu

### Debug Information

Enable debug logging by checking the browser console and terminal output for detailed information about:
- Panel conversion status
- System information
- Error messages
- Event handling

## Architecture

This example demonstrates best practices for:

1. **Service Architecture**: Clean separation between UI and backend logic
2. **Error Handling**: Proper error handling for cross-platform compatibility
3. **Event Management**: Proper window event handling and cleanup
4. **UI/UX**: Modern, responsive interface design
5. **Platform Integration**: Native system tray and panel behaviors

## Further Reading

- [Wails v3 Documentation](https://v3.wails.io/)
- [macOS NSPanel Documentation](https://developer.apple.com/documentation/appkit/nspanel)
- [System Tray Examples](../systray-basic/)

## Contributing

This example serves as a reference implementation for NSPanel integration. Feel free to:
- Report issues
- Suggest improvements
- Submit pull requests
- Use as a base for your own applications