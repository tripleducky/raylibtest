package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1200
	screenHeight = 600
)

var (
	running  = true
	bkgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc  rl.Rectangle
	playerDest rl.Rectangle

	playerX, playerY float64
	playerC          float64
	playerXPos       float64
	playerYPos       float64
	playerSpeed      float64 = 3

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D
)

func main() {

	for running {
		input()
		update()
		render()
	}
	quit()
}

func drawScene() {
	rl.DrawFPS(20, 20)
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

}

func input() {
	//calculates player following the mouse
	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		playerX = (float64((float32(rl.GetMouseX()) + 48) - playerDest.X))
		playerY = (float64((float32(rl.GetMouseY()) + 48) - playerDest.Y))
		playerC = math.Sqrt(playerX*playerX + playerY*playerY)
		playerXPos = playerSpeed * (playerX / playerC)
		playerYPos = playerSpeed * (playerY / playerC)

		playerDest.X = float32(playerXPos) + playerDest.X
		playerDest.Y = float32(playerYPos) + playerDest.Y
	}
	if rl.IsKeyPressed(rl.KeyP) {
		musicPaused = !musicPaused
	}
}
func update() {
	running = !rl.WindowShouldClose()
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
	var assAss = rl.NewVector2(playerDest.X, playerDest.Y)
	cam.Target = assAss
}
func render() {

	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)
	drawScene()

	rl.EndMode2D()

	rl.EndDrawing()
}

func init() {

	rl.InitWindow(screenWidth, screenHeight, "Game Test")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("res/Characters/BasicCharakterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(screenWidth/2, screenHeight/2, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/Averyfarm.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X), float32(playerDest.Y)), 0, 1)
}
func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
