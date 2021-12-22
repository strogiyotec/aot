
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public final class Day22Part1 {

    public static void main(String[] args) throws Exception {
        final Pattern pattern = Pattern.compile("(on|off) x=([-0-9]+)..([^,]+),y=([-0-9]+)..([^,]+),z=([-0-9]+)..(.*)");
        final String[] lines = Files
                .readString(Path.of("/home/strogiyotec/IdeaProjects/GoLang/aoc/aot/2021/day22/input.txt")).split("\n");
        final List<Input> collect = Stream.of(lines).map(line -> new Input(line, pattern))
                .filter(line -> line.fromX >= -50 && line.fromX <= 50).collect(Collectors.toList());
        final Set<Tuple> tuples = new HashSet<>(collect.size() << 1);
        for (final Input input : collect) {
            for (int i = input.fromX; i <= input.toX; i++) {
                for (int j = input.fromY; j <= input.toY; j++) {
                    for (int k = input.fromZ; k <= input.toZ; k++) {
                        if (input.on) {
                            tuples.add(new Tuple(i, j, k));
                        } else {
                            tuples.remove(new Tuple(i, j, k));
                        }
                    }
                }
            }
        }
        System.out.println(tuples.size());
    }

    private static class Tuple {
        final int first;
        final int second;
        final int third;

        public Tuple(final int first, final int second, final int third) {
            this.first = first;
            this.second = second;
            this.third = third;
        }

        @Override
        public boolean equals(final Object o) {
            if (this == o)
                return true;
            if (o == null || this.getClass() != o.getClass())
                return false;
            final Tuple tuple = (Tuple) o;
            return this.first == tuple.first && this.second == tuple.second && this.third == tuple.third;
        }

        @Override
        public int hashCode() {
            return Objects.hash(this.first, this.second, this.third);
        }
    }

    private static class Input {
        private final int fromX;
        private final int toX;
        private final int fromY;
        private final int toY;
        private final int fromZ;
        private final int toZ;
        private final boolean on;

        public Input(final String line, final Pattern pattern) {
            final Matcher matcher = pattern.matcher(line);
            if (matcher.find()) {
                this.on = "on".equals(matcher.group(1));

                this.fromX = Integer.parseInt(matcher.group(2));
                this.toX = Integer.parseInt(matcher.group(3));

                this.fromY = Integer.parseInt(matcher.group(4));
                this.toY = Integer.parseInt(matcher.group(5));

                this.fromZ = Integer.parseInt(matcher.group(6));
                this.toZ = Integer.parseInt(matcher.group(7));
            } else {
                throw new IllegalArgumentException("Wrong input");
            }
        }
    }
}
