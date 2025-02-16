.PHONY: wrapper
wrapper:
	g++ -c sum/src/sum.cpp -o sum/sum.o
	g++ -c sum/src/wrapper.cpp -o sum/wrapper.o
	command ar rcs sum/libwrapper.a sum/sum.o sum/wrapper.o
	rm sum/sum.o sum/wrapper.o