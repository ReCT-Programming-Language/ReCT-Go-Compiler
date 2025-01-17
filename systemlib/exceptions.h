#ifndef EXCEPTIONS_H
#define EXCEPTIONS_H

// define the throwing function (very athletic)
// ============================================

#ifdef __cplusplus
extern "C" {
#endif

// standard exception throwing
void exc_Throw(char *message);

// exception shortcuts
void exc_ThrowIfNull(void *pointer);
void exc_ThrowIfInvalidCast(class_Any* from, Standard_vTable *to, const char *toFingerprint);

#ifdef __cplusplus
}
#endif

#endif