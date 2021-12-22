import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.regex.Pattern;
import lombok.EqualsAndHashCode;

public final class Day22Part2 {

    private static Pattern commandPtn = Pattern.compile(".*?x=(-?\\d+)\\.\\.(-?\\d+)\\,y=(-?\\d+)\\.\\.(-?\\d+)\\,z=(-?\\d+)\\.\\.(-?\\d+)");

    public static void main(String[] args) throws IOException {
        final String[] lines = Files.readString(
            Path.of(
                "input.txt"
            )
        ).split("\n");
        var cuboids = new ArrayList<Cuboid>(16);
        for (final String line : lines) {
            var current = Cuboid.parse(line);
            var intersections = new ArrayList<Cuboid>();
            for (final var prev : cuboids) {
                if (current.intersect(prev) ) {
                    intersections.add(current.intersection(prev));
                }
            }
            cuboids.addAll(intersections);
            if (current.on) {
                cuboids.add(current);
            }
        }
        long answer = 0L;
        for (final Cuboid cuboid : cuboids) {
            if (cuboid.on) {
                answer += cuboid.volume();
            } else {
                answer -= cuboid.volume();
            }
        }
        System.out.println(answer);
    }

    @EqualsAndHashCode
    static class Cube {
        final int x, y, z;

        Cube(int x, int y, int z) {
            this.x = x;
            this.y = y;
            this.z = z;
        }
    }

    @EqualsAndHashCode
    static class Cuboid {
        final int x1, x2, y1, y2, z1, z2;
        final boolean on;

        Cuboid(int x1, int x2, int y1, int y2, int z1, int z2, boolean on) {
            this.x1 = Math.min(x1, x2);
            this.x2 = Math.max(x1, x2);
            this.y1 = Math.min(y1, y2);
            this.y2 = Math.max(y1, y2);
            this.z1 = Math.min(z1, z2);
            this.z2 = Math.max(z1, z2);
            this.on = on;
        }

        static Cuboid parse(String s) {
            var m = commandPtn.matcher(s);
            m.matches();
            return new Cuboid(
                Integer.parseInt(m.group(1)),
                Integer.parseInt(m.group(2)),
                Integer.parseInt(m.group(3)),
                Integer.parseInt(m.group(4)),
                Integer.parseInt(m.group(5)),
                Integer.parseInt(m.group(6)),
                s.startsWith("on")
            );
        }

        Cuboid intersection(Cuboid prev) {
            return new Cuboid(
                Math.max(this.x1, prev.x1),
                Math.min(this.x2, prev.x2),
                Math.max(this.y1, prev.y1),
                Math.min(this.y2, prev.y2),
                Math.max(this.z1, prev.z1),
                Math.min(this.z2, prev.z2),
                !prev.on
            );
        }

        boolean intersect(Cuboid other) {
            return Math.min(this.x2, other.x2) >= Math.max(this.x1, other.x1) &&
                Math.min(this.y2, other.y2) >= Math.max(this.y1, other.y1) &&
                Math.min(this.z2, other.z2) >= Math.max(this.z1, other.z1);
        }

        long volume() {
            return ((long) (x2 - x1 + 1)) * (y2 - y1 + 1) * (z2 - z1 + 1);
        }
    }
}

