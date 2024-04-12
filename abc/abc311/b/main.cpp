#include <cstdio>
using namespace std;
const int N=1000003;
char s[N];
bool a[N];
int main(){
    int n,m;
    scanf("%d%d",&n,&m);
    for(int i=1;i<=n;++i){
        scanf("%s",s+1);
        for(int j=1;j<=m;++j) if(s[j]=='x') a[j]=1;
    }
    int res=0;
    for(int l=1;l<=m;++l)
        for(int r=l;r<=m;++r){
            if(a[r]) break;
            if(r-l+1>res) res=r-l+1;
        }
    printf("%d\n",res);
    return 0;
}