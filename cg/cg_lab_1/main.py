import glfw
from OpenGL.GL import *
from OpenGL.GLU import *
import numpy
import math

delta = 0.0
angle = 0.0
posx = 0.0
posy = 0.0
size = 0.0
radius=0.25
sides=100

def main():
	if not glfw.init():
		return
	window = glfw.create_window(640, 480, "Hello World", None, None)
	if not window:
		glfw.terminate()
		return

	glfw.make_context_current(window)
	glfw.set_key_callback(window, key_callback)
	glfw.set_scroll_callback(window, scroll_callback)


	while not glfw.window_should_close(window):
		display_circle(window)
		glfw.swap_buffers(window)
		glfw.poll_events()
		
	glfw.destroy_window(window)
	glfw.terminate()

def display(window):
	global angle
	glClear(GL_COLOR_BUFFER_BIT)
	glLoadIdentity()
	glClearColor(1.0, 1.0, 1.0, 1.0)
	glPushMatrix()
	glRotatef(angle, 0, 0, 1)
	glBegin(GL_POLYGON)
	glColor3f(0.1,0.1,0.1)
	
	glVertex2f(posx + size + 0.5,posy + size + 0.5)
	glColor3f(0.35,0.0,0.89)
	glVertex2f(posx - size + -0.5,posy + size + 0.5)
	glColor3f(0.0,1.0,1.0)
	
	glVertex2f(posx - size,posy - size)
	
	glVertex2f(posx - size + -0.5,posy - size + -0.5)
	glColor3f(0.78,0.23,1.0)
	glVertex2f(posx + size + 0.5,posy - size + -0.5)
	glEnd()
	glPopMatrix()
	angle += delta

def display_circle(window):
	glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)    
	glLoadIdentity()
	glClearColor(1.0, 1.0, 1.0, 1.0)
	glPushMatrix()
	global angle
	glRotatef(angle, 0, 0, 1)
	glBegin(GL_POLYGON)   
	global sides
	global redius 
	for i in range(100):    
		cosine= radius * math.cos(i*2*math.pi/sides) + posx    
		sine  = radius * math.sin(i*2*math.pi/sides) + posy    
		glVertex2f(cosine,sine)
		glColor3f((0.78 + i/10) % 1 ,(0.23 + i /10)%1 ,(1.0 + i / 10)%1)
	glEnd()
	glPopMatrix()
	angle += delta

def key_callback(window, key, scancode, action, mods):
	global delta
	global angle
	global size
	global posx
	global posy
	if action == glfw.PRESS or glfw.REPEAT:
		if key == glfw.KEY_RIGHT:
			delta = -0.1
		if key == glfw.KEY_LEFT: # glfw.KEY_LEFT
			delta = 0.1
		if key == glfw.KEY_UP: # glfw.KEY_LEFT
			posx += 0.01
		if key == glfw.KEY_DOWN: # glfw.KEY_LEFT
			posx -= 0.01
		if key == glfw.KEY_BACKSPACE:
			delta = 0.0

def scroll_callback(window, xoffset, yoffset):
	global size
	if (xoffset > 0):
		size -= yoffset/10
	else:
		size += yoffset/10

if __name__ == "__main__":
	main()
