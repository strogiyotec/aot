
public final class Day21Part1 {

    public static void main(String[] args) {
        final Player firstPlayer = new Player(7, 1000);
        final Player secondPlayer = new Player(8, 1000);
        final Dice dice = new Dice();
        while (true) {
            firstPlayer.add(dice.roll());
            if (firstPlayer.isWin()) {
                System.out.println(secondPlayer.score * dice.times);
                break;
            }
            secondPlayer.add(dice.roll());
            if (secondPlayer.isWin()) {
                System.out.println(firstPlayer.score * dice.times);
                break;
            }
        }
    }

    private static class Dice {
        private int times;

        private int val = 1;

        public int roll() {
            int score = 0;
            for (int i = 0; i < 3; i++) {
                score += this.val;
                this.val++;
            }
            this.times += 3;
            return score;
        }
    }

    private static class Player {

        private int startingPosition;

        private int limit;

        private int score = 0;

        public Player(final int starting, final int limit) {
            this.startingPosition = starting;
            this.limit = limit;
        }

        public void add(final int score) {
            this.startingPosition += score;
            this.startingPosition = (this.startingPosition - 1) % 10 + 1;
            this.score += this.startingPosition;
        }

        public boolean isWin() {
            return this.score >= this.limit;
        }

        @Override
        public String toString() {
            return "Player{" + "score=" + this.score + ",position=" + this.startingPosition + '}';
        }
    }
}
