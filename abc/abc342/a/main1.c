#include<stdio.h>
int main(){
    int n,i,b[310]={};
    char s[110],a;
    scanf("%s",s);
    for(i=0;s[i];i++)b[s[i]]++;
    for(i=0;i<300;i++){
        if(b[i]==1)a=i;
    }
    for(i=0;s[i]!=a;i++);
    printf("%d\n",i+1);
    return 0;
}