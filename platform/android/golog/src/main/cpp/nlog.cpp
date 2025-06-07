#include <jni.h>
#include <string>
//#include "libgolog.h"
#include <libgolog.h>

extern "C" JNIEXPORT jstring JNICALL
Java_com_example_golog_NativeLib_stringFromJNI(
        JNIEnv* env,
        jobject /* this */) {
    std::string hello = "Hello from C++";
    return env->NewStringUTF(hello.c_str());
}

char* jstringToCString(JNIEnv* env, jstring jstr) {
    const char* c=  env->GetStringUTFChars(jstr, nullptr);
    return const_cast<char*>(c);
}

extern "C" JNIEXPORT void JNICALL
Java_com_example_golog_NativeLib_initLog(
        JNIEnv* env,
        jobject thiz,
        jstring path) {
    char* logDir = jstringToCString(env, path);
    InitLogger(logDir);
}

extern "C" JNIEXPORT void JNICALL
Java_com_example_golog_NativeLib_closeLogger(
        JNIEnv* env,
        jobject thiz) {
    CloseLogger();
}

extern "C" JNIEXPORT void JNICALL
Java_com_example_golog_NativeLib_logDebug(
        JNIEnv* env,
        jobject thiz,
        jstring tag,
        jstring msg
        ) {
    char* ctag = jstringToCString(env, tag);
    char* cmsg = jstringToCString(env, msg);
    LogDebug(ctag, cmsg);
}

extern "C" JNIEXPORT void JNICALL
Java_com_example_golog_NativeLib_logInfo(
        JNIEnv* env,
        jobject thiz,
        jstring tag,
        jstring msg
        ) {
    char* ctag = jstringToCString(env, tag);
    char* cmsg = jstringToCString(env, msg);
    LogInfo(ctag, cmsg);
}

extern "C" JNIEXPORT void JNICALL
Java_com_example_golog_NativeLib_logWarn(
        JNIEnv* env,
        jobject thiz,
        jstring tag,
        jstring msg
        ) {
    char* ctag = jstringToCString(env, tag);
    char* cmsg = jstringToCString(env, msg);
    LogWarn(ctag, cmsg);
}

extern "C" JNIEXPORT void JNICALL
Java_com_example_golog_NativeLib_logError(
        JNIEnv* env,
        jobject thiz,
        jstring tag,
        jstring msg
        ) {
    char* ctag = jstringToCString(env, tag);
    char* cmsg = jstringToCString(env, msg);
    LogError(ctag, cmsg);
}