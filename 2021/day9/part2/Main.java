import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Comparator;
import java.util.LinkedList;
import java.util.List;
import java.util.PriorityQueue;
import java.util.Queue;

public class Main {
    public static void main(final String[] args) throws Exception {
        var lines = Files.readString(Path.of("/home/strogiyotec/IdeaProjects/GoLang/aoc/aot/2021/day9/part2/input.txt"))
                .split("\n");
        final int[][] matrix = new int[lines.length][lines[0].length()];
        for (int i = 0; i < matrix.length; i++) {
            var row = lines[i];
            for (int j = 0; j < matrix[0].length; j++) {
                matrix[i][j] = Character.getNumericValue(row.charAt(j));
            }
        }
        final List<int[]> list = new ArrayList<>();
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[0].length; j++) {
                int currentValue = matrix[i][j];
                boolean left = isLower(currentValue, i, j - 1, matrix);
                boolean right = isLower(currentValue, i, j + 1, matrix);
                boolean down = isLower(currentValue, i - 1, j, matrix);
                boolean up = isLower(currentValue, i + 1, j, matrix);
                if (left && right && down && up) {
                    list.add(new int[] { i, j });
                }
            }
        }
        System.out.println(getHighestBasins(list, matrix));
    }

    static int getHighestBasins(final List<int[]> list, final int[][] matrix) {
        final PriorityQueue<Integer> max = new PriorityQueue(Comparator.reverseOrder());
        final boolean[][] cache = new boolean[matrix.length][matrix[0].length];
        final int[][] directions = { { 0, 1 }, { 0, -1 }, { -1, 0 }, { 1, 0 }, };
        for (final int[] pair : list) {
            final Queue<int[]> queue = new LinkedList<>();
            queue.add(pair);
            int cnt = 0;
            while (!queue.isEmpty()) {
                final int[] current = queue.poll();
                final int row = current[0];
                final int col = current[1];
                final int currentVal = matrix[row][col];
                if (!cache[row][col]) {
                    cnt++;
                    cache[row][col] = true;
                    for (final int[] direction : directions) {
                        final int nextRow = direction[0] + row;
                        final int nextCol = direction[1] + col;
                        if (nextRow >= 0 && nextRow < matrix.length && nextCol >= 0 && nextCol < matrix[0].length
                                && !cache[nextRow][nextCol]) {
                            final int nextVal = matrix[nextRow][nextCol];
                            if (nextVal > currentVal && nextVal != 9) {
                                queue.add(new int[] { nextRow, nextCol });
                            }
                        }
                    }
                }
            }
            max.add(cnt);
        }
        return max.poll() * max.poll() * max.poll();
    }

    static boolean isLower(int value, int nextRow, int nextCol, int[][] matrix) {
        if (nextRow < 0 || nextRow >= matrix.length || nextCol < 0 || nextCol >= matrix[0].length) {
            return true;
        }
        return value < matrix[nextRow][nextCol];
    }
}
