
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

void bindcpu(int i)
{
        cpu_set_t mask;
        CPU_ZERO(&mask);
        CPU_SET(i, &mask);
        if (pthread_setaffinity_np(pthread_self(), sizeof(mask), &mask) < 0)
        {
                fprintf(stderr, "set thread affinity failed\n");
                exit(-1);
        }
}


int main(){
       // auto a = new pad;
        auto a = new nopad;
        a->a = 0;
        a->b = 0;
        std::thread ta([&]{
                bindcpu(0);
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
                bindcpu(1);
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
