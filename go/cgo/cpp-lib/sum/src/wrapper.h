#ifndef WRAPPER_H
#define WRAPPER_H

#ifdef __cplusplus
extern "C"
{
#endif

    typedef void *SumWrapper;
    SumWrapper Init();
    void Destroy(SumWrapper s);
    int Sum(SumWrapper s, int a, int b);

#ifdef __cplusplus
}
#endif
#endif