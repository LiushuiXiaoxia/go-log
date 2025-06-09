ANDROID_OUT=build_output/jniLibs
ANDROID_SDK=$(ANDROID_HOME)
NDK_BIN=$(ANDROID_SDK)/ndk/27.0.12077973/toolchains/llvm/prebuilt/darwin-x86_64/bin
GO_FILES=golog.go
SO_NAME=libgolog.so

android: android-armv7a android-arm64 android-x86 android-x86_64
	rm -rf platform/android/golog/src/main/jniLibs
	cp -R build_output/jniLibs platform/android/golog/src/main/jniLibs

android-armv7a:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	CC=$(NDK_BIN)/armv7a-linux-androideabi21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/armeabi-v7a/$(SO_NAME) $(GO_FILES)

android-arm64:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm64 \
	CC=$(NDK_BIN)/aarch64-linux-android21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/arm64-v8a/$(SO_NAME) $(GO_FILES)

android-x86:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=386 \
	CC=$(NDK_BIN)/i686-linux-android21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/x86/$(SO_NAME) $(GO_FILES)

android-x86_64:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=amd64 \
	CC=$(NDK_BIN)/x86_64-linux-android21-clang \
	go build -buildmode=c-shared -o $(ANDROID_OUT)/x86_64/$(SO_NAME) $(GO_FILES)

