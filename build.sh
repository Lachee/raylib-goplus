echo "======= Converting Header Files"
cd raylib-convert
go run main.go

echo "======= Copying Generated Files"
echo "Copying Audio"
cp out/audio_gen.go ../raylib/audio_gen.go


echo "======= Building Library"
echo " ( this might take a while, please wait ) "
cd ../raylib
go build .


echo "======= Finished"