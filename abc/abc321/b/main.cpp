#include<bits/stdc++.h>

using namespace std;

int main(){
  int n,x;
  cin >> n >> x;
  vector<int> a(n-1);
  for(auto &nx : a){cin >> nx;}

  int l=0,r=100;
  while(l<=r){
    int te=(l+r)/2;
    vector<int> b=a;
    b.push_back(te);
    int sum=0,ma=0,mi=100;
    for(auto &nx : b){
      sum+=nx;
      ma=max(ma,nx);
      mi=min(mi,nx);
    }
    int score=sum-ma-mi;
    if(score>=x){r=te-1;}
    else{l=te+1;}
  }
  if(l>100){l=-1;}
  cout << l << "\n";
  return 0;
}
