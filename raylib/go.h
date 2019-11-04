static char**makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}
static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}
static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++) free(a[i]);
				free(a);
}
