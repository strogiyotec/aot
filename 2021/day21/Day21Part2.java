
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

public final class Day21Part2 {

    public static void main(String[] args) {
        final long[] dp = dp(7, 0L, 8, 0L, true, new HashMap<>());
        System.out.println(Arrays.toString(dp));
        System.out.println(Math.max(dp[0], dp[1]));
    }

    private static long[] dp(final int firstPosition, final long firstScore, final int secondPosition,
            final long secondScore, final boolean firstTurn, final Map<Tuple, long[]> cache) {
        final Tuple tuple = new Tuple(firstPosition, firstScore, secondPosition, secondScore, firstTurn);
        if (cache.containsKey(tuple)) {
            return cache.get(tuple);
        }
        if (firstScore >= 21) {
            return new long[] { 1, 0 };
        }
        if (secondScore >= 21) {
            return new long[] { 0, 1 };
        }
        long w1 = 0;
        long w2 = 0;
        for (int i = 1; i <= 3; i++) {
            for (int j = 1; j <= 3; j++) {
                for (int k = 1; k <= 3; k++) {
                    if (firstTurn) {
                        int nextFirstPosition = i + j + k + firstPosition;
                        nextFirstPosition = (nextFirstPosition - 1) % 10 + 1;
                        final long nextFirstScore = firstScore + nextFirstPosition;
                        final long[] dp = dp(nextFirstPosition, nextFirstScore, secondPosition, secondScore, false,
                                cache);
                        w1 += dp[0];
                        w2 += dp[1];
                    } else {
                        int nextSecondPosition = i + j + k + secondPosition;
                        nextSecondPosition = (nextSecondPosition - 1) % 10 + 1;
                        final long nextSecondScore = secondScore + nextSecondPosition;
                        final long[] dp = dp(firstPosition, firstScore, nextSecondPosition, nextSecondScore, true,
                                cache);
                        w1 += dp[0];
                        w2 += dp[1];
                    }
                }
            }
        }
        cache.put(tuple, new long[] { w1, w2 });
        return cache.get(tuple);
    }

    private static class Tuple {
        private final int first;
        private final long second;
        private final int third;
        private final long fourth;
        private final boolean turn;

        public Tuple(final int first, final long second, final int third, final long fourth, final boolean turn) {
            this.first = first;
            this.second = second;
            this.third = third;
            this.fourth = fourth;
            this.turn = turn;
        }

        @Override
        public boolean equals(final Object o) {
            if (this == o)
                return true;
            if (o == null || this.getClass() != o.getClass())
                return false;
            final Tuple tuple = (Tuple) o;
            return this.first == tuple.first && this.second == tuple.second && this.third == tuple.third
                    && this.fourth == tuple.fourth && this.turn == tuple.turn;
        }

        @Override
        public int hashCode() {
            return Objects.hash(this.first, this.second, this.third, this.fourth, this.turn);
        }
    }

}
