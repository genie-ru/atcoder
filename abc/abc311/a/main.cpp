#include<bits/stdc++.h> 
using namespace std;
char a,b,c,s;
int l;
int main(){
	cin>>l;
	for(int i=0;i<l;++i){
		cin>>s;
		if(s=='A')a=1;
		if(s=='B')b=1;
		if(s=='C')c=1;
		if(a&&b&&c){
			cout<<i+1;
			return 0;
		}
	}
}