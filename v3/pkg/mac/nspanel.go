//go:build darwin

// Package mac provides NSPanel functionality for macOS
package mac

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AppKit

#import <Foundation/Foundation.h>
#import <AppKit/AppKit.h>

// NSPanel creation and management functions
void* createNSPanel(void* nsWindow);
void setPanelDelegate(void* nsPanel, void* delegate);
void setPanelReleasedWhenClosed(void* nsPanel, bool released);
void* getNSPanelFromWindow(void* nsWindow);
bool isNSPanel(void* nsWindow);
void setPanelLevel(void* nsPanel, int level);
void setPanelFloating(void* nsPanel, bool floating);
void setPanelHidesOnDeactivate(void* nsPanel, bool hidesOnDeactivate);
void setPanelBecomesKeyOnlyIfNeeded(void* nsPanel, bool becomesKeyOnlyIfNeeded);
void setPanelCanBecomeKey(void* nsPanel, bool canBecomeKey);
void setPanelWorksWhenModal(void* nsPanel, bool worksWhenModal);

// Create a new panel delegate
void* createPanelDelegate();

// NSPanel implementation
void* createNSPanel(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if (!window) {
        return NULL;
    }
    
    // Get the window's properties
    NSRect frame = [window frame];
    NSWindowStyleMask styleMask = [window styleMask];
    NSBackingStoreType backing = [window backingType];
    
    // Create new NSPanel with same properties
    NSPanel* panel = [[NSPanel alloc] 
        initWithContentRect:[window contentRectForFrameRect:frame]
        styleMask:styleMask
        backing:backing
        defer:NO];
    
    // Copy properties from the original window
    [panel setContentView:[window contentView]];
    [panel setTitle:[window title]];
    [panel setDelegate:[window delegate]];
    [panel setBackgroundColor:[window backgroundColor]];
    [panel setOpaque:[window isOpaque]];
    [panel setHasShadow:[window hasShadow]];
    [panel setAlphaValue:[window alphaValue]];
    [panel setLevel:[window level]];
    [panel setCollectionBehavior:[window collectionBehavior]];
    
    // Set default panel behaviors
    [panel setFloatingPanel:YES];
    [panel setHidesOnDeactivate:YES];
    [panel setBecomesKeyOnlyIfNeeded:YES];
    
    return panel;
}

void setPanelDelegate(void* nsPanel, void* delegate) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel && delegate) {
        [panel setDelegate:(id<NSWindowDelegate>)delegate];
    }
}

void setPanelReleasedWhenClosed(void* nsPanel, bool released) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        [panel setReleasedWhenClosed:released];
    }
}

void* getNSPanelFromWindow(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    if ([window isKindOfClass:[NSPanel class]]) {
        return window;
    }
    return NULL;
}

bool isNSPanel(void* nsWindow) {
    NSWindow* window = (NSWindow*)nsWindow;
    return [window isKindOfClass:[NSPanel class]];
}

void setPanelLevel(void* nsPanel, int level) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        [panel setLevel:level];
    }
}

void setPanelFloating(void* nsPanel, bool floating) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        [panel setFloatingPanel:floating];
    }
}

void setPanelHidesOnDeactivate(void* nsPanel, bool hidesOnDeactivate) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        [panel setHidesOnDeactivate:hidesOnDeactivate];
    }
}

void setPanelBecomesKeyOnlyIfNeeded(void* nsPanel, bool becomesKeyOnlyIfNeeded) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        [panel setBecomesKeyOnlyIfNeeded:becomesKeyOnlyIfNeeded];
    }
}

void setPanelCanBecomeKey(void* nsPanel, bool canBecomeKey) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        // This would require subclassing NSPanel to override canBecomeKeyWindow
        // For now, we'll use setBecomesKeyOnlyIfNeeded as an alternative
        [panel setBecomesKeyOnlyIfNeeded:!canBecomeKey];
    }
}

void setPanelWorksWhenModal(void* nsPanel, bool worksWhenModal) {
    NSPanel* panel = (NSPanel*)nsPanel;
    if (panel) {
        [panel setWorksWhenModal:worksWhenModal];
    }
}

// Simple panel delegate for basic event handling
@interface SimplePanelDelegate : NSObject <NSWindowDelegate>
@end

@implementation SimplePanelDelegate

- (void)windowDidBecomeKey:(NSNotification *)notification {
    // Panel became key window
}

- (void)windowDidResignKey:(NSNotification *)notification {
    // Panel resigned key window
}

- (void)windowWillClose:(NSNotification *)notification {
    // Panel will close
}

@end

// Create a new panel delegate
void* createPanelDelegate() {
    SimplePanelDelegate* delegate = [[SimplePanelDelegate alloc] init];
    return delegate;
}

*/
import "C"
import (
	"unsafe"
)

// NSPanel represents a macOS NSPanel
type NSPanel struct {
	ptr unsafe.Pointer
}

// PanelDelegateCallback is the callback function type for panel delegate events
type PanelDelegateCallback func(methodName string, panel *NSPanel)

// PanelDelegate represents an NSPanel delegate
type PanelDelegate struct {
	ptr      unsafe.Pointer
	callback PanelDelegateCallback
}

// CreateNSPanel creates a new NSPanel from an existing NSWindow
func CreateNSPanel(nsWindow unsafe.Pointer) *NSPanel {
	panel := C.createNSPanel(nsWindow)
	if panel == nil {
		return nil
	}
	return &NSPanel{ptr: panel}
}

// GetNSPanelFromWindow returns an NSPanel if the window is a panel, nil otherwise
func GetNSPanelFromWindow(nsWindow unsafe.Pointer) *NSPanel {
	panel := C.getNSPanelFromWindow(nsWindow)
	if panel == nil {
		return nil
	}
	return &NSPanel{ptr: panel}
}

// IsNSPanel checks if the given window is an NSPanel
func IsNSPanel(nsWindow unsafe.Pointer) bool {
	return bool(C.isNSPanel(nsWindow))
}

// SetReleasedWhenClosed sets whether the panel is released when closed
func (p *NSPanel) SetReleasedWhenClosed(released bool) {
	C.setPanelReleasedWhenClosed(p.ptr, C.bool(released))
}

// SetLevel sets the window level of the panel
func (p *NSPanel) SetLevel(level int) {
	C.setPanelLevel(p.ptr, C.int(level))
}

// SetFloating sets whether the panel is floating
func (p *NSPanel) SetFloating(floating bool) {
	C.setPanelFloating(p.ptr, C.bool(floating))
}

// SetHidesOnDeactivate sets whether the panel hides when it becomes inactive
func (p *NSPanel) SetHidesOnDeactivate(hides bool) {
	C.setPanelHidesOnDeactivate(p.ptr, C.bool(hides))
}

// SetBecomesKeyOnlyIfNeeded sets whether the panel becomes key only if needed
func (p *NSPanel) SetBecomesKeyOnlyIfNeeded(onlyIfNeeded bool) {
	C.setPanelBecomesKeyOnlyIfNeeded(p.ptr, C.bool(onlyIfNeeded))
}

// SetCanBecomeKey sets whether the panel can become the key window
func (p *NSPanel) SetCanBecomeKey(canBecomeKey bool) {
	C.setPanelCanBecomeKey(p.ptr, C.bool(canBecomeKey))
}

// SetWorksWhenModal sets whether the panel works when a modal window is present
func (p *NSPanel) SetWorksWhenModal(worksWhenModal bool) {
	C.setPanelWorksWhenModal(p.ptr, C.bool(worksWhenModal))
}

// SetDelegate sets the delegate for the panel
func (p *NSPanel) SetDelegate(delegate *PanelDelegate) {
	if delegate != nil {
		C.setPanelDelegate(p.ptr, delegate.ptr)
	}
}

// GetNativePointer returns the native pointer to the NSPanel
func (p *NSPanel) GetNativePointer() unsafe.Pointer {
	return p.ptr
}

// CreatePanelDelegate creates a new panel delegate
func CreatePanelDelegate() *PanelDelegate {
	delegate := C.createPanelDelegate()
	if delegate == nil {
		return nil
	}
	
	return &PanelDelegate{
		ptr: delegate,
	}
}

// Common NSWindowLevel constants for convenience
const (
	NSNormalWindowLevel       = 0
	NSFloatingWindowLevel     = 3
	NSModalPanelWindowLevel   = 8
	NSPopUpMenuWindowLevel    = 101
	NSScreenSaverWindowLevel  = 1000
	NSStatusWindowLevel       = 25
	NSTornOffMenuWindowLevel  = 3
	NSMainMenuWindowLevel     = 24
)