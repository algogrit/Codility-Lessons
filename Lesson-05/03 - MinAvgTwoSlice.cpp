int solution(vector<int> &A) {

  // Find prefix sum.
  int N = A.size();
  vector<int> ps(N + 1, 0);

  for (int i = 1; i <= N; i++) {
      ps[i] = A[i - 1] + ps[i - 1];
  }

  int lft_idx, min_lft_idx;
  double avg_here, min_avg, avg_of_two, avg_with_prev;

  // Initialize variables at the first possible slice (A[0:1]).
  lft_idx = min_lft_idx = 0;
  avg_here = min_avg = (A[0] + A[1]) / 2.0;

  // Find min average of every slice that ends at ith element,
  // starting at i = 2.
  for (int i = 2; i < N; i ++) {

      // average of A[lft_idx : i]
      avg_with_prev = ((double) ps[i + 1] - ps[lft_idx]) /
                      (i - lft_idx + 1);

      // average of A[i - 1 : i]
      avg_of_two = (A[i - 1] + A[i]) / 2.0;

      // Find minimum and update lft_idx of slice
      // (previous lft_idx or i - 1).
      if (avg_of_two < avg_with_prev) {
          avg_here = avg_of_two;
          lft_idx = i - 1;
      }
      else
          avg_here = avg_with_prev;

      // Keep track of minimum so far and its left index.
      if (avg_here < min_avg) {
          min_avg = avg_here;
          min_lft_idx = lft_idx;
      }
  }

  return min_lft_idx;
}
