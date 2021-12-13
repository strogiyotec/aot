import java.nio.file.Files;
import java.nio.file.Path;

public class Main {

    public static void main(String[] args) throws Exception {
        var lines = Files.readString(Path.of("input.txt")).split("\n");
        final int[][] matrix = new int[lines.length][lines[0].length()];
        for (int i = 0; i < matrix.length; i++) {
            var row = lines[i];
            for (int j = 0; j < matrix[0].length; j++) {
                matrix[i][j] = Character.getNumericValue(row.charAt(j));
            }
        }
        int cnt = 0;
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[0].length; j++) {
                int currentValue = matrix[i][j];
                boolean left = isLower(currentValue, i, j - 1, matrix);
                boolean right = isLower(currentValue, i, j + 1, matrix);
                boolean down = isLower(currentValue, i - 1, j, matrix);
                boolean up = isLower(currentValue, i + 1, j, matrix);
                if (left && right && down && up) {
                    cnt += matrix[i][j] + 1;
                }
            }
        }
        System.out.println(cnt);
    }

    static boolean isLower(int value, int nextRow, int nextCol, int[][] matrix) {
        if (nextRow < 0 || nextRow >= matrix.length || nextCol < 0 || nextCol >= matrix[0].length) {
            return true;
        }
        return value < matrix[nextRow][nextCol];
    }

}
