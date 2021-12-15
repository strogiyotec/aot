import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

public class Day13 {

    public static void main(String[] args) throws Exception {
        final String[] lines = Files.readString(Path.of("input.txt")).split("\n");
        final MatrixWithSpace matrixWithSpace = createMatrix(lines);
        boolean[][] matrix = matrixWithSpace.matrix;
        final List<FoldAlong> folds = folds(lines, matrixWithSpace.spaceCoordinate);
        for (final FoldAlong fold : folds) {
            if (fold.alongX) {
                for (int row = 0; row < matrix.length; row++) {
                    for (int column = fold.distance; column < matrix[0].length; column++) {
                        if (matrix[row][column]) {
                            matrix[row][column] = false;
                            final int foldedCol = column - (column - fold.distance << 1);
                            matrix[row][foldedCol] = true;
                        }
                    }
                }
                final boolean[][] temp = new boolean[matrix.length][fold.distance];
                for (int i = 0; i < temp.length; i++) {
                    System.arraycopy(matrix[i], 0, temp[i], 0, temp[0].length);
                }
                matrix = temp;
            } else {
                for (int row = fold.distance; row < matrix.length; row++) {
                    for (int column = 0; column < matrix[0].length; column++) {
                        if (matrix[row][column]) {
                            matrix[row][column] = false;
                            final int foldedRow = row - (row - fold.distance << 1);
                            matrix[foldedRow][column] = true;
                        }
                    }
                }
                final boolean[][] temp = new boolean[fold.distance][matrix[0].length];
                for (int i = 0; i < temp.length; i++) {
                    System.arraycopy(matrix[i], 0, temp[i], 0, temp[0].length);
                }
                matrix = temp;
            }
        }
        for (final boolean[] booleans : matrix) {
            for (int j = 0; j < matrix[0].length; j++) {
                if (booleans[j]) {
                    System.out.print('#');
                } else {
                    System.out.print('.');
                }
            }
            System.out.println();
        }
    }

    private static List<FoldAlong> folds(final String[] lines, final int spaceCoordinate) {
        final List<FoldAlong> folds = new ArrayList<>(lines.length - spaceCoordinate + 1);
        for (int i = spaceCoordinate + 1; i < lines.length; i++) {
            folds.add(new FoldAlong(lines[i]));
        }
        return folds;
    }

    private static MatrixWithSpace createMatrix(final String[] lines) {
        final Collection<int[]> coordinatesList = new ArrayList<>(lines.length);
        int lengthRow = 0;
        int lengthCol = 0;
        int spaceIndex = 0;
        for (int i = 0; i < lines.length; i++) {
            final String line = lines[i];
            if (line.isBlank()) {
                spaceIndex = i;
                break;
            }
            final String[] coordinates = line.split(",");
            final int col = Integer.parseInt(coordinates[0]);
            final int row = Integer.parseInt(coordinates[1]);
            coordinatesList.add(new int[] { row, col });
            lengthRow = Math.max(lengthRow, row);
            lengthCol = Math.max(lengthCol, col);
        }
        final boolean[][] matrix = new boolean[lengthRow + 1][lengthCol + 1];
        for (final int[] coordinate : coordinatesList) {
            matrix[coordinate[0]][coordinate[1]] = true;
        }
        return new MatrixWithSpace(spaceIndex, matrix);
    }

    static class MatrixWithSpace {
        int spaceCoordinate;
        boolean[][] matrix;

        public MatrixWithSpace(final int spaceCoordinate, final boolean[][] matrix) {
            this.spaceCoordinate = spaceCoordinate;
            this.matrix = matrix;
        }
    }

    static class FoldAlong {
        boolean alongX;
        int distance;

        public FoldAlong(final String line) {
            final String[] parts = line.split(" ");
            final String last = parts[2];
            this.alongX = last.charAt(0) == 'x';
            this.distance = Integer.parseInt(last.split("=")[1]);
        }
    }
}
