import 'dart:ffi';
import 'dart:io';

import 'generated_bindings.dart';

void main(List<String> arguments) {
	
	final nativeAddLib = DynamicLibrary.open("${Directory.current.path}/libmain.so");
	final int Function(int x, int y) nativeAdd = nativeAddLib
    .lookup<NativeFunction<Int32 Function(Int32, Int32)>>('native_add')
    .asFunction();
}
