package main

import (
	_ "embed"
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/icons"
)

//go:embed assets
var assets embed.FS

// AppService provides methods to interact with the app
type AppService struct {
	app    *application.App
	window *application.WebviewWindow
}

// NewAppService creates a new AppService
func NewAppService(app *application.App) *AppService {
	return &AppService{
		app: app,
	}
}

func (a *AppService) SetWindow(window *application.WebviewWindow) {
	a.window = window
}

// ConvertToPanel converts the window to an NSPanel
func (a *AppService) ConvertToPanel() error {
	if a.window == nil {
		return fmt.Errorf("window not set")
	}
	
	err := a.window.ConvertToPanel()
	if err != nil {
		return fmt.Errorf("failed to convert to panel: %w", err)
	}
	
	// Configure panel behavior
	_ = a.window.SetPanelFloating(true)
	_ = a.window.SetPanelHidesOnDeactivate(true)
	_ = a.window.SetPanelBecomesKeyOnlyIfNeeded(true)
	_ = a.window.SetPanelWorksWhenModal(false)
	_ = a.window.SetPanelReleasedWhenClosed(true)
	
	return nil
}

// IsPanel returns whether the window is currently a panel
func (a *AppService) IsPanel() bool {
	if a.window == nil {
		return false
	}
	return a.window.IsPanel()
}

// ShowWindow shows the window
func (a *AppService) ShowWindow() {
	if a.window != nil {
		a.window.Show()
		a.window.Focus()
	}
}

// HideWindow hides the window
func (a *AppService) HideWindow() {
	if a.window != nil {
		a.window.Hide()
	}
}

// TogglePanel toggles the panel visibility
func (a *AppService) TogglePanel() {
	if a.window == nil {
		return
	}
	
	if a.window.IsVisible() {
		a.window.Hide()
	} else {
		a.window.Show()
		a.window.Focus()
	}
}

// GetSystemInfo returns information about the system
func (a *AppService) GetSystemInfo() map[string]interface{} {
	return map[string]interface{}{
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
		"version": runtime.Version(),
		"isPanel": a.IsPanel(),
	}
}

func main() {
	// Create the application
	app := application.New(application.Options{
		Name:        "NSPanel + System Tray Demo",
		Description: "A demonstration of NSPanel functionality with system tray integration",
		Services: []application.Service{
			application.NewService(&AppService{}),
		},
		Assets: application.AssetOptions{
			FS: assets,
		},
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})

	// Create the main window (initially hidden)
	window := app.NewWebviewWindow().
		SetTitle("NSPanel Demo").
		SetSize(400, 600).
		SetResizable(false).
		SetFrameless(false).
		SetAlwaysOnTop(true).
		SetURL("/").
		SetHidden(true)

	// Create system tray
	systemTray := app.NewSystemTray()

	// Set the tray icon based on platform
	if runtime.GOOS == "darwin" {
		systemTray.SetTemplateIcon(icons.SystrayMacTemplate)
	} else {
		systemTray.SetIcon(icons.WailsIcon)
	}

	// Attach the window to the system tray
	systemTray.AttachWindow(window).WindowOffset(10)

	// Set up window event handlers
	window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		// Prevent window from closing, just hide it instead
		e.Cancel()
		window.Hide()
	})

	// Get the app service and set the window reference
	appService := app.GetService(&AppService{}).(*AppService)
	appService.SetWindow(window)

	// On macOS, automatically convert to panel when window is ready
	if runtime.GOOS == "darwin" {
		window.OnWindowEvent(events.Mac.WebViewDidFinishNavigation, func(_ *application.WindowEvent) {
			// Convert to panel after a short delay to ensure everything is ready
			go func() {
				// Wait a bit for the window to be fully initialized
				select {
				case <-context.Background().Done():
					return
				default:
					err := appService.ConvertToPanel()
					if err != nil {
						log.Printf("Failed to convert to panel: %v", err)
					} else {
						log.Println("Successfully converted window to NSPanel")
					}
				}
			}()
		})
	}

	// Create context menu for system tray
	menu := app.NewMenu().
		Add("Show Panel").OnClick(func(ctx context.Context, data application.MenuItemClickEventData) {
			appService.ShowWindow()
		}).
		AddSeparator().
		Add("Toggle Panel").OnClick(func(ctx context.Context, data application.MenuItemClickEventData) {
			appService.TogglePanel()
		}).
		AddSeparator().
		Add("Quit").OnClick(func(ctx context.Context, data application.MenuItemClickEventData) {
			app.Quit()
		})

	systemTray.SetMenu(menu)

	// Set tooltip
	systemTray.SetTooltip("NSPanel Demo - Click to toggle panel")

	// Run the application
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}