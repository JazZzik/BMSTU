import glfw
from OpenGL.GL import *
from OpenGL.GLU import *
import numpy
import math
from math import *


sc = 0.2
an = 0


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
	     
	mz = [1, 0, 0, 0,
		  0, 1, 0, 0,
		  0, 0, -1, 0,
		  0, 0, 0, 1]
	mx = [0, 0, -1, 0,
		  0, 1, 0, 0,
		  -1, 0, 0, 0,
		  0, 0, 0, 1]
	my = [1, 0, 0, -0,
		  0, 0, -1, 0,
		  0, -1, 0, 0,
		  0, 0, 0, 1]
	
		
	glEnable(GL_DEPTH_TEST)
	glfw.set_key_callback(window, key_callback)
	glfw.set_scroll_callback(window, scroll_callback)


	while not glfw.window_should_close(window):
		glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
		global an
		glMatrixMode(GL_PROJECTION)
		glLoadIdentity()
		glMultMatrixd(m1)
		glMatrixMode(GL_MODELVIEW)
		glLoadIdentity()
		display_cube(window, 0, 0)

		glMatrixMode(GL_PROJECTION)
		glLoadIdentity()
		glMultMatrixd(mx)
		glMatrixMode(GL_MODELVIEW)
		glLoadIdentity()
		glTranslate(0.2, 0.2, -0.4)
		display_cube(window, an, sc)

		glMatrixMode(GL_PROJECTION)
		glLoadIdentity()
		glMultMatrixd(mz)
		glMatrixMode(GL_MODELVIEW)
		glLoadIdentity()
		glTranslate(0.4, -0.4, -0.0)
		display_cube(window, an, sc)

		glMatrixMode(GL_PROJECTION)
		glLoadIdentity()
		glMultMatrixd(my)
		glMatrixMode(GL_MODELVIEW)
		glLoadIdentity()
		glTranslate(0, 0.4, -0.4)
		display_cube(window, an, sc)



		glfw.swap_buffers(window)
		glfw.poll_events()
		
	glfw.destroy_window(window)
	glfw.terminate()

def display_cube(window, angle, scale):
	glScale(0.3, 0.3, 0.3)

	glRotatef(angle, 0, 1, 0)
	glRotatef(angle, 1.0, 0.0, 0.0)
	glRotatef(angle, 0.0, 1.0, 0.0)
	glRotatef(angle, 0.0, 0.0, 1.0)
	glBegin(GL_POLYGON)
	glColor3f(0.0, 0.3, 0.3)  
	glVertex3f(0.4, -0.4, -0.4)
	glVertex3f(0.4, 0.4, -0.4)
	glVertex3f(-0.4, 0.4, -0.4)
	glVertex3f(-0.4, -0.4, -0.4)
	glEnd()
	
	glBegin(GL_POLYGON)
	glColor3f(0.5, 0.0, 0.5) 
	glVertex3f(0.4, -0.4, 0.4)
	glVertex3f(0.4, 0.4, 0.4)
	glVertex3f(-0.4, 0.4, 0.4)
	glVertex3f(-0.4, -0.4, 0.4)
	glEnd()

	glBegin(GL_POLYGON)
	glColor3f(1.0, 0.0, 1.0)   
	glVertex3f(0.4, -0.4, -0.4)
	glVertex3f(0.4, 0.4, -0.4)
	glVertex3f(0.4, 0.4, 0.4)
	glVertex3f(0.4, -0.4, 0.4)
	glEnd()	

	glBegin(GL_POLYGON)
	glColor3f(0.0, 1.0, 0.0)   
	glVertex3f(-0.4, -0.4, 0.4)
	glVertex3f(-0.4, 0.4, 0.4)
	glVertex3f(-0.4, 0.4, -0.4)
	glVertex3f(-0.4, -0.4, -0.4)
	glEnd()	

	glBegin(GL_POLYGON)
	glColor3f(0.0, 0.0, 1.0)   
	glVertex3f(0.4, 0.4, 0.4)
	glVertex3f(0.4, 0.4, -0.4)
	glVertex3f(-0.4, 0.4, -0.4)
	glVertex3f(-0.4, 0.4, 0.4)
	glEnd()

	glBegin(GL_POLYGON)
	glColor3f(1.0, 0.0, 0.0)   
	glVertex3f(0.4, -0.4, -0.4)
	glVertex3f(0.4, -0.4, 0.4)
	glVertex3f(-0.4, -0.4, 0.4)
	glVertex3f(-0.4, -0.4, -0.4)
	glEnd()

def key_callback(window, key, scancode, action, mods):
	global an
	if action == glfw.PRESS or glfw.REPEAT:
		if glfw.get_key(window, glfw.KEY_L):
			glPolygonMode(GL_FRONT_AND_BACK, GL_LINE)
		if glfw.get_key(window, glfw.KEY_O):
			glPolygonMode(GL_FRONT_AND_BACK, GL_FILL)
		if key == glfw.KEY_RIGHT:
			an += 10
		if key == glfw.KEY_LEFT:
			an -= 10

def scroll_callback(window, xoffset, yoffset):
	global sc
	if (xoffset > 0):
		sc -= yoffset/10
	else:
		sc += yoffset/10


if __name__ == "__main__":
	main()
