
#include <cstdio>
#include <unistd.h>
#include <thread>
#include <atomic>

// command line: gcc main.cpp -std=c++11 -o main -lstdc++ -lpthread

struct nopad {
        uint64_t a;
        uint64_t b;
};

struct pad {
        uint64_t a;
        uint64_t pa[8];
        uint64_t pb[8];
        uint64_t b;
};



int main(){
       // auto a = new pad;
        auto a = new nopad;
        a->a = 0;
        a->b = 0;
        std::thread ta([&]{
                while(1){
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                        a->a++;
                };
        });

        std::thread tb([&]{
                while(1){
                        a->b++;
                        a->b++;
                        a->b++;
                        a->b++;
                        a->b++;
                        a->b++;
                        a->b++;
                        a->b++;
                        a->b++;
                }
        });

        sleep(1*60);

        printf("%lld\n",a->a+a->b);
        return 0;
}
