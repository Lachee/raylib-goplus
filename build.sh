echo "======= Converting Header Files"
cd raylib-convert
go run main.go

echo "======= Copying Generated Files"
echo "GUI";  	cp out/raygui_gen.go ../raylib/raygui_gen.go
echo "Audio";  	cp out/audio_gen.go ../raylib/audio_gen.go
echo "Texture";  cp out/texture_gen.go ../raylib/texture_gen.go
echo "Camera";  cp out/camera_gen.go ../raylib/camera_gen.go
echo "Gestures";  cp out/gestures_gen.go ../raylib/gestures_gen.go
echo "Input";  cp out/input_gen.go ../raylib/input_gen.go
echo "Shader";  cp out/shader_gen.go ../raylib/shader_gen.go


echo "======= Building Library"
echo " ( this might take a while, please wait ) "
cd ../raylib

if [[ $(go build .) ]];
then
	echo "======= Finished Unsucesfully"
else
	echo "======= Reporting Lint Issues"
	#golint .
	echo "======= Finished Succesfully"
fi
