#pragma once

#include <X11/Xlib.h>
#include <X11/XKBlib.h>
#include <X11/Xutil.h>
#include <X11/extensions/Xrandr.h>
#include <X11/extensions/XTest.h>
#include <X11/extensions/Xfixes.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

// for computing xrandr modelines at runtime
#include <libxcvt/libxcvt.h>

extern void goCreateScreenSize(int index, int width, int height, int mwidth, int mheight);
extern void goSetScreenRates(int index, int rate_index, short rate);

Display *getXDisplay(void);
int XDisplayOpen(char *input);
void XDisplayClose(void);

void XMove(int x, int y);
void XCursorPosition(int *x, int *y);
void XScroll(int deltaX, int deltaY);
void XButton(unsigned int button, int down);

typedef struct xkeyentry_t {
  KeySym keysym;
  KeyCode keycode;
  struct xkeyentry_t *next;
} xkeyentry_t;

static void XKeyEntryAdd(KeySym keysym, KeyCode keycode);
static KeyCode XKeyEntryGet(KeySym keysym);
static KeyCode XkbKeysymToKeycode(Display *dpy, KeySym keysym);
void XKey(KeySym keysym, int down);

Status XSetScreenConfiguration(int width, int height, short rate);
void XGetScreenConfiguration(int *width, int *height, short *rate);
void XGetScreenConfigurations();
void XCreateScreenMode(int width, int height, short rate);
XRRModeInfo XCreateScreenModeInfo(int hdisplay, int vdisplay, short vrefresh);

void XSetKeyboardModifier(unsigned char mod, int on);
unsigned char XGetKeyboardModifiers();
XFixesCursorImage *XGetCursorImage(void);

char *XGetScreenshot(int *w, int *h);
