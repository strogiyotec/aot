import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;

public final class Day11 {

    static int cnt = 0;

    public static void main(String[] args) throws IOException {
        final int[][] matrix = createMatrix();
        for (int i = 0; i < Integer.MAX_VALUE; i++) {
            final boolean[][] cache = new boolean[matrix.length][matrix[0].length];
            cnt = 0;
            for (int row = 0; row < matrix.length; row++) {
                for (int column = 0; column < matrix[0].length; column++) {
                    if (!cache[row][column]) {
                        if (matrix[row][column] == 9) {
                            flash(row, column, matrix, cache);
                            if (cnt == matrix.length * matrix[0].length) {
                                System.out.println(i + 1);
                                return;
                            }
                        } else {
                            matrix[row][column]++;
                        }
                    }
                }
            }
        }
        System.out.println(Arrays.deepToString(matrix));
    }

    private static int[][] createMatrix() throws IOException {
        final String[] lines = Files.readString(Path.of("input.txt")).split("\n");
        final int[][] matrix = new int[lines.length][lines[0].length()];
        for (int i = 0; i < matrix.length; i++) {
            var row = lines[i];
            for (int j = 0; j < matrix[0].length; j++) {
                matrix[i][j] = Character.getNumericValue(row.charAt(j));
            }
        }
        return matrix;
    }

    private static void flash(final int row, final int column, final int[][] matrix, final boolean[][] cache) {
        if (row < 0 || row >= matrix.length || column < 0 || column >= matrix[0].length || cache[row][column]) {
            return;
        }
        if (matrix[row][column] == 9) {
            cache[row][column] = true;
            cnt++;
            matrix[row][column] = 0;
            flash(row + 1, column, matrix, cache);
            flash(row - 1, column, matrix, cache);
            flash(row, column - 1, matrix, cache);
            flash(row, column + 1, matrix, cache);
            flash(row + 1, column + 1, matrix, cache);
            flash(row + 1, column - 1, matrix, cache);
            flash(row - 1, column + 1, matrix, cache);
            flash(row - 1, column - 1, matrix, cache);
        } else {
            matrix[row][column]++;
        }
    }
}
