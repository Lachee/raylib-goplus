echo "======= Converting Header Files"
cd raylib-convert
go run main.go

echo "======= Copying Generated Files"
echo "GUI";  	cp out/raygui_gen.go ../raylib/raygui_gen.go
echo "Audio";  	cp out/audio_gen.go ../raylib/audio_gen.go
echo "Camera";  cp out/camera_gen.go ../raylib/camera_gen.go


echo "======= Building Library"
echo " ( this might take a while, please wait ) "
cd ../raylib
go build .

echo "======= Reporting Lint Issues"
golint .


echo "======= Finished"