
public final class Day17 {

    public static void main(String[] args) {
        final int fromX = 241;
        final int toX = 273;
        final int fromY = -97;
        final int toY = -63;
        int cnt = 0;
        for (int x = 1; x < toX + 1; x++) {
            for (int y = fromY; y < Math.abs(fromY); y++) {
                int decreaseXStep = x;
                int decreaseYStep = y;
                int currentXPoint = 0;
                int currentYPoint = 0;
                int maxYPath = 0;
                for (int time = 1; time < 2 * Math.abs(fromY) + 1; time++) {
                    currentXPoint += decreaseXStep;
                    currentYPoint += decreaseYStep;
                    //rules
                    decreaseXStep = Math.max(decreaseXStep - 1, 0);
                    decreaseYStep--;
                    if (currentXPoint >= fromX && currentXPoint <= toX && currentYPoint >= fromY && currentYPoint <= toY) {
                        cnt++;
                        break;
                    } else if (currentXPoint > toX || currentYPoint < fromY) {
                        break;
                    }
                }
            }
        }
        System.out.println(cnt);
    }
}

