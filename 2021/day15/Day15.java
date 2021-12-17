
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Comparator;
import java.util.HashSet;
import java.util.PriorityQueue;
import java.util.Set;

public final class Day15 {

    public static void main(String[] args) throws Exception {
        final String[] lines = Files.readString(Path.of("input.txt")).split("\n");
        final int[][] matrix = asMatrix(lines);
        final int[][] increasedMatrix = increasedMatrix(matrix);
        bfs(increasedMatrix);
    }

    private static void bfs(final int[][] matrix) {
        final PriorityQueue<int[]> queue = new PriorityQueue<>(Comparator.comparingInt(value -> value[0]));
        queue.add(new int[] { 0, 0, 0 });
        final int[][] directions = { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 }, };
        final Set<String> visited = new HashSet<>(matrix.length);
        while (!queue.isEmpty()) {
            final int[] poll = queue.poll();
            final int cost = poll[0];
            final int row = poll[1];
            final int col = poll[2];
            final String line = row + "-" + col;
            if (visited.contains(line)) {
                continue;
            }
            visited.add(line);
            if (row == matrix.length - 1 && col == matrix[0].length - 1) {
                System.out.println(cost);
                return;
            }
            for (final int[] direction : directions) {
                final int nextRow = direction[0] + row;
                final int nextCol = direction[1] + col;
                if (invalidRange(matrix, nextRow, nextCol)) {
                    continue;
                }
                queue.add(new int[] { cost + matrix[nextRow][nextCol], nextRow, nextCol });
            }
        }
        System.out.println("No");
    }

    private static boolean invalidRange(final int[][] matrix, final int nextRow, final int nextCol) {
        return nextRow < 0 || nextRow >= matrix.length || nextCol < 0 || nextCol >= matrix[0].length;
    }

    private static int[][] increasedMatrix(final int[][] origin) {
        final int[][] increasedMatrix = new int[origin.length * 5][origin[0].length * 5];
        final int width = origin[0].length;
        final int height = origin.length;
        for (int i = 0; i < origin.length; i++) {
            for (int j = 0; j < origin[0].length; j++) {
                for (int offsetRow = 0; offsetRow < 5; offsetRow++) {
                    for (int offsetCol = 0; offsetCol < 5; offsetCol++) {
                        final int nextRow = offsetRow * height + i;
                        final int nextCol = offsetCol * width + j;
                        final int nextValue;
                        final int currentValue = origin[i][j] + offsetRow + offsetCol;
                        if (currentValue > 9) {
                            nextValue = currentValue - 9;
                        } else {
                            nextValue = currentValue;
                        }
                        increasedMatrix[nextRow][nextCol] = nextValue;
                    }
                }
            }
        }
        return increasedMatrix;
    }

    private static int[][] asMatrix(final String[] lines) {
        final int[][] matrix = new int[lines.length][lines[0].length()];
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[0].length; j++) {
                matrix[i][j] = Character.getNumericValue(lines[i].charAt(j));
            }
        }
        return matrix;
    }
}
