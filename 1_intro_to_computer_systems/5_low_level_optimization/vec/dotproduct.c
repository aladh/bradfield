#include "vec.h"


data_t dotproduct(vec_ptr u, vec_ptr v) {
   data_t sum0 = 0, sum1 = 0, sum2 = 0, sum3 = 0, *u_data, *v_data, length, i, limit;

   length = vec_length(u);
   u_data = u->data;
   v_data = v->data;
   limit = length - 3;

   for (i = 0; i < limit; i+=4) { // we can assume both vectors are same length
      sum0 += u_data[i] * v_data[i];
      sum1 += u_data[i+1] * v_data[i+1];
      sum2 += u_data[i+2] * v_data[i+2];
      sum3 += u_data[i+3] * v_data[i+3];
   }

   for (;i < length; i++) {
      sum0 += u_data[i] * v_data[i];
   }

   return sum0 + sum1 + sum2 + sum3;
}
