#include <iostream>

using namespace std;

int main()
{
    int t, a[100];
    cin >> t;

    for(int i=0; i<t; i++)
        cin>> a[i];
    for(int i=0; i<t; i++)
        cout << a[i];
    return 0;
}
