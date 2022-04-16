import 'dart:ffi';

import 'generated_bindings.dart';

void main(List<String> arguments) {
	NativeLibrary library = NativeLibrary(DynamicLibrary.open('libmain.so'));
	print(library.HelloWorld());
}
