#ifndef GO_HELPER
#define GO_HELPER

static char**makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}
static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}
static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++) free(a[i]);
				free(a);
}


// Security check in case no GRAPHICS_API_OPENGL_* defined
#if !defined(GRAPHICS_API_OPENGL_11) && \
    !defined(GRAPHICS_API_OPENGL_21) && \
    !defined(GRAPHICS_API_OPENGL_33) && \
    !defined(GRAPHICS_API_OPENGL_ES2)
        #define GRAPHICS_API_OPENGL_33
#endif

// Security check in case multiple GRAPHICS_API_OPENGL_* defined
#if defined(GRAPHICS_API_OPENGL_11)
    #if defined(GRAPHICS_API_OPENGL_21)
        #undef GRAPHICS_API_OPENGL_21
    #endif
    #if defined(GRAPHICS_API_OPENGL_33)
        #undef GRAPHICS_API_OPENGL_33
    #endif
    #if defined(GRAPHICS_API_OPENGL_ES2)
        #undef GRAPHICS_API_OPENGL_ES2
    #endif
#endif

#if defined(GRAPHICS_API_OPENGL_21)
    #define GRAPHICS_API_OPENGL_33
#endif



#if defined(GRAPHICS_API_OPENGL_11)
    #if defined(__APPLE__)
        #include <OpenGL/gl.h>      // OpenGL 1.1 library for OSX
        #include <OpenGL/glext.h>
    #else
        #include <GL/gl.h>              // OpenGL 1.1 library
    #endif
#endif

#if defined(GRAPHICS_API_OPENGL_21)
    #define GRAPHICS_API_OPENGL_33      // OpenGL 2.1 uses mostly OpenGL 3.3 Core functionality
#endif

#if defined(GRAPHICS_API_OPENGL_33)
    #if defined(__APPLE__)
        #include <OpenGL/gl3.h>         // OpenGL 3 library for OSX
        #include <OpenGL/gl3ext.h>      // OpenGL 3 extensions library for OSX
    #else
          #include "external/glad.h"  // GLAD extensions loading library, includes OpenGL headers
    #endif
#endif

#if defined(GRAPHICS_API_OPENGL_ES2)
    #include <EGL/egl.h>                // EGL library
    #include <GLES2/gl2.h>              // OpenGL ES 2.0 library
    #include <GLES2/gl2ext.h>           // OpenGL ES 2.0 extensions library
#endif

static void reportGlErrors() {
  GLenum err;
  while((err = glGetError()) != GL_NO_ERROR)
  {
      TraceLog(LOG_ERROR, "[GL ERROR] %02X", err);
  }
}

#endif
