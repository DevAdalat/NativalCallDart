import 'dart:ffi';
import 'dart:io';

import 'generated_bindings.dart';

void main(List<String> arguments) {

	NativeLibrary library = NativeLibrary(DynamicLibrary.open('${Directory.current.path}/libmain.so'));
	print(library.helloWorld());
	print(library.getUid());
}
