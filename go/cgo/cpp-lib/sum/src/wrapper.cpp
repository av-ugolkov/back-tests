// num.cpp
#include "sum.h"
#include "wrapper.h"

extern "C"
{
    SumWrapper Init()
    {
        cxxSum *ret = new cxxSum();
        return (void *)ret;
    }

    void Destroy(SumWrapper n)
    {
        cxxSum *num = (cxxSum *)n;
        delete num;
    }

    int Sum(SumWrapper n, int a, int b)
    {
        cxxSum *num = (cxxSum *)n;
        return num->Sum(a, b);
    }
}