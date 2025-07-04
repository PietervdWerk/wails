# Wails v3 NSPanel Integration

## Overview

This document summarizes the complete NSPanel integration added to Wails v3, inspired by the Tauri NSPanel plugin. This integration allows macOS applications to convert regular windows to NSPanels, providing better system integration and user experience for utility-style applications.

## Implementation Summary

### 1. Core NSPanel Package (`v3/pkg/mac/nspanel.go`)

A new package providing low-level NSPanel functionality:

**Features:**
- Convert NSWindow to NSPanel
- Panel behavior configuration (floating, hiding, key management)
- Proper resource management
- Native macOS integration

**Key Functions:**
```go
// Panel creation and conversion
func CreateNSPanel(nsWindow unsafe.Pointer) *NSPanel
func GetNSPanelFromWindow(nsWindow unsafe.Pointer) *NSPanel
func IsNSPanel(nsWindow unsafe.Pointer) bool

// Panel configuration
func (p *NSPanel) SetFloating(floating bool)
func (p *NSPanel) SetHidesOnDeactivate(hides bool)
func (p *NSPanel) SetBecomesKeyOnlyIfNeeded(onlyIfNeeded bool)
func (p *NSPanel) SetWorksWhenModal(worksWhenModal bool)
func (p *NSPanel) SetReleasedWhenClosed(released bool)
```

### 2. WebviewWindow NSPanel Integration (`v3/pkg/application/webview_window_darwin.go`)

Extended the macOS WebviewWindow implementation with NSPanel support:

**Added C Functions:**
- `convertWindowToPanel()` - Core window-to-panel conversion
- `windowIsPanel()` - Panel type checking  
- `getWindowAsPanel()` - Panel pointer retrieval
- Panel configuration functions

**Added Go Methods:**
```go
func (w *macosWebviewWindow) ConvertToPanel() error
func (w *macosWebviewWindow) IsPanel() bool
func (w *macosWebviewWindow) SetPanelFloating(floating bool) error
func (w *macosWebviewWindow) SetPanelHidesOnDeactivate(hides bool) error
func (w *macosWebviewWindow) SetPanelBecomesKeyOnlyIfNeeded(onlyIfNeeded bool) error
func (w *macosWebviewWindow) SetPanelWorksWhenModal(worksWhenModal bool) error
func (w *macosWebviewWindow) SetPanelReleasedWhenClosed(released bool) error
```

### 3. Public WebviewWindow API (`v3/pkg/application/webview_window.go`)

Added public methods to the WebviewWindow interface:

**Cross-Platform Methods:**
```go
func (w *WebviewWindow) ConvertToPanel() error
func (w *WebviewWindow) IsPanel() bool
func (w *WebviewWindow) SetPanelFloating(floating bool) error
func (w *WebviewWindow) SetPanelHidesOnDeactivate(hides bool) error
func (w *WebviewWindow) SetPanelBecomesKeyOnlyIfNeeded(onlyIfNeeded bool) error
func (w *WebviewWindow) SetPanelWorksWhenModal(worksWhenModal bool) error
func (w *WebviewWindow) SetPanelReleasedWhenClosed(released bool) error
```

All methods include platform detection and return appropriate errors on non-macOS systems.

### 4. Complete Example Application (`v3/examples/nspanel-systray/`)

A comprehensive demonstration app featuring:

**Backend Features:**
- System tray integration
- Automatic NSPanel conversion on macOS
- Window management through service layer
- Cross-platform compatibility

**Frontend Features:**
- Modern, responsive UI with glassmorphism effects
- Real-time status monitoring
- Platform-specific feature detection
- Interactive controls for all NSPanel functions
- Toast notifications for user feedback

**Files Structure:**
```
v3/examples/nspanel-systray/
├── main.go                 # Main application logic
├── assets/index.html       # Frontend UI
├── go.mod                  # Go module definition
└── README.md              # Comprehensive documentation
```

## Key Features Implemented

### 1. Panel Conversion
- **Seamless Conversion**: Convert any WebviewWindow to NSPanel while preserving all properties
- **Property Preservation**: Maintains window content, title, delegate, styling, and behaviors
- **Resource Management**: Proper cleanup of original window resources

### 2. Panel Behaviors
- **Floating**: Panels float above other windows
- **Auto-Hide**: Hide when becoming inactive (user clicks elsewhere)
- **Key Management**: Become key window only when needed
- **Modal Interaction**: Configure behavior when modal windows are present
- **Memory Management**: Automatic or manual resource release when closed

### 3. System Integration
- **System Tray**: Full integration with system tray for panel management
- **Context Menu**: Right-click menu with panel controls
- **Native Appearance**: Uses native macOS panel appearance and behaviors
- **Window Management**: Proper integration with macOS window management

### 4. Cross-Platform Support
- **Platform Detection**: Automatically detects macOS vs other platforms
- **Graceful Degradation**: NSPanel features disabled on non-macOS systems
- **Error Handling**: Appropriate error messages for unsupported platforms
- **Unified API**: Same API works across platforms with appropriate behavior

## Technical Implementation Details

### Architecture
The implementation follows a layered architecture:

1. **C Layer**: Direct NSPanel and Objective-C integration
2. **CGO Bridge**: Safe bridging between C and Go
3. **Platform Layer**: macOS-specific implementation
4. **Cross-Platform Layer**: Unified API with platform detection
5. **Application Layer**: High-level window management

### Memory Management
- Automatic reference counting for NSPanel objects
- Proper cleanup when windows are destroyed
- Optional manual resource release configuration
- Thread-safe operations using Wails' main thread dispatcher

### Error Handling
- Platform-appropriate error messages
- Graceful fallbacks for unsupported features
- Comprehensive error propagation through the API stack
- User-friendly error reporting in the example app

## Benefits Over Regular Windows

### User Experience
1. **Better Focus Management**: Panels integrate better with macOS focus system
2. **Utility Window Behavior**: Designed for tool and utility windows
3. **Space Efficiency**: Can auto-hide to conserve screen space
4. **Native Feel**: Behaves like native macOS utility panels

### Developer Experience
1. **Simple API**: Easy conversion with single method call
2. **Fine-Grained Control**: Configure all panel behaviors individually
3. **Cross-Platform**: Same code works on all platforms
4. **Well-Documented**: Comprehensive documentation and examples

### System Integration
1. **Window Management**: Better integration with Spaces and Mission Control
2. **Performance**: More efficient than regular windows for utility purposes
3. **Accessibility**: Proper accessibility integration
4. **Resource Usage**: Optimized memory and resource usage

## Usage Examples

### Basic Panel Conversion
```go
// Create window
window := app.NewWebviewWindow()

// Convert to panel on macOS
err := window.ConvertToPanel()
if err != nil {
    log.Printf("Panel conversion failed: %v", err)
}
```

### Full Panel Configuration
```go
// Convert and configure panel
if err := window.ConvertToPanel(); err == nil {
    window.SetPanelFloating(true)
    window.SetPanelHidesOnDeactivate(true)
    window.SetPanelBecomesKeyOnlyIfNeeded(true)
    window.SetPanelWorksWhenModal(false)
    window.SetPanelReleasedWhenClosed(true)
}
```

### System Tray Integration
```go
// Create system tray
systemTray := app.NewSystemTray()
systemTray.AttachWindow(window)

// Convert to panel for better tray integration
window.ConvertToPanel()
```

## Testing & Validation

### Automated Testing
- Platform detection tests
- API surface validation
- Error handling verification
- Cross-platform compatibility

### Manual Testing
- Panel conversion on macOS
- Behavior verification (floating, hiding, etc.)
- System tray integration
- Resource cleanup verification
- Cross-platform graceful degradation

## Comparison with Tauri NSPanel Plugin

### Similarities
- Core NSPanel functionality
- Panel behavior configuration
- Resource management
- Cross-platform API design

### Improvements
- **Integrated Design**: Native integration with Wails architecture
- **Type Safety**: Go's type system provides better safety than raw Tauri APIs
- **Error Handling**: More comprehensive error handling and propagation
- **Documentation**: More thorough documentation and examples
- **System Integration**: Better integration with Wails' window management

### Wails-Specific Benefits
- **Service Integration**: Works seamlessly with Wails service architecture
- **Event System**: Integrates with Wails' event system
- **Asset Management**: Works with Wails' asset serving
- **Development Experience**: Consistent with other Wails APIs

## Future Enhancements

### Potential Improvements
1. **Panel Delegate Events**: Full delegate event system for panel lifecycle
2. **Panel Styling**: Extended styling options for panels
3. **Multiple Panel Types**: Support for different NSPanel subtypes
4. **Animation Support**: Panel show/hide animations
5. **Advanced Positioning**: Smart positioning relative to system tray

### Community Contributions
- Additional panel behaviors
- Platform-specific optimizations
- Enhanced documentation
- More example applications
- Performance improvements

## Conclusion

This NSPanel integration provides a comprehensive, well-designed addition to Wails v3 that:

1. **Enables Native macOS Integration**: Full NSPanel support with all behaviors
2. **Maintains Cross-Platform Compatibility**: Graceful handling of all platforms  
3. **Follows Wails Design Principles**: Consistent with existing Wails architecture
4. **Provides Complete Examples**: Production-ready example application
5. **Ensures Developer Experience**: Simple APIs with comprehensive documentation

The implementation successfully bridges the gap between web-based UI and native macOS window management, providing developers with powerful tools for creating professional utility applications while maintaining the ease of use that Wails is known for.