import glfw
from OpenGL.GL import *
from OpenGL.GLU import *
import numpy
import math
from math import *


sc = 1
an = 0
n = 30
a = 0.5
b = 0.8

def main():
    if not glfw.init():
        return
    window = glfw.create_window(900, 900, "Square", None, None)
    if not window:
        glfw.terminate()
        return

    glfw.make_context_current(window)
    teta = asin(0.625/sqrt(2))
    fi = asin (0.625/sqrt(2 - 0.625**2))
    m1 = [cos(fi), sin(teta)*sin(fi), sin(fi) * cos(teta), 0,
         0, cos(teta), -sin(teta), 0,
         sin(fi), -cos(fi)*sin(teta), -cos(fi)*cos(teta), 0,
         0, 0, 0, 1]
    
        
    glEnable(GL_DEPTH_TEST)
    glfw.set_key_callback(window, key_callback)
    glfw.set_scroll_callback(window, scroll_callback)
    display = (400, 300)
    glMatrixMode(GL_PROJECTION)
    gluPerspective(45, (display[0]/display[1]), 0.1, 50.0)
    glMatrixMode(GL_MODELVIEW)
    gluLookAt(0, -8, 0, 0, 0, 0, 0, 0, 1)
    viewMatrix = glGetFloatv(GL_MODELVIEW_MATRIX)
    glLoadIdentity()

    while not glfw.window_should_close(window):
        glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
        global an, sc
        glMatrixMode(GL_PROJECTION)
        glLoadIdentity()
        glMultMatrixd(m1)
        glMatrixMode(GL_MODELVIEW)
        glLoadIdentity()
        display_cube(window, an, sc)


        glfw.swap_buffers(window)
        glfw.poll_events()
        
    glfw.destroy_window(window)
    glfw.terminate()


def display_cube(window, angle, scale):
    glRotatef(angle, 1.0, 0.0, 0.0)
    glRotatef(angle, 0.0, 1.0, 0.0)
    glRotatef(angle, 0.0, 0.0, 1.0)
    sphere = gluNewQuadric() #Create new sphere
    glTranslatef(-1.5, 0, 0) #Move to the place
    glColor4f(0.5, 0.2, 0.2, 1) #Put color
    gluSphere(sphere, sc, 32, 16) #Draw sphere


def key_callback(window, key, scancode, action, mods):
    global an, sc
    if action == glfw.PRESS or glfw.REPEAT:
        if glfw.get_key(window, glfw.KEY_L):
            glPolygonMode(GL_FRONT_AND_BACK, GL_LINE)
        if glfw.get_key(window, glfw.KEY_O):
            glPolygonMode(GL_FRONT_AND_BACK, GL_FILL)
        if key == glfw.KEY_RIGHT:
            an += 10
        if key == glfw.KEY_LEFT:
            an -= 10
        if key == glfw.KEY_UP:
            sc += 0.1
        if key == glfw.KEY_DOWN:
            sc -= 0.1

def scroll_callback(window, xoffset, yoffset):
    global sc
    if (xoffset > 0):
        sc -= yoffset/10
    else:
        sc += yoffset/10


if __name__ == "__main__":
    main()
