import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Set;

public final class Part1 {

    static int cnt;

    public static void main(String[] args) throws IOException {
        final String[] lines = Files.readString(Path.of("input.txt")).split("\n");
        final Map<Vertex, Set<Vertex>> graph = new HashMap<>(lines.length << 1);
        for (final String line : lines) {
            final String[] parts = line.split("-");
            final Vertex first = new Vertex(parts[0]);
            final Vertex second = new Vertex(parts[1]);
            graph.putIfAbsent(first, new HashSet<>());
            graph.putIfAbsent(second, new HashSet<>());
            if (parts[0].equals("start")) {
                graph.get(first).add(second);
            } else if (parts[1].equals("start")) {
                graph.get(second).add(first);
            } else {
                graph.get(second).add(first);
                graph.get(first).add(second);
            }
        }
        Vertex start = new Vertex("start");
        for (final Vertex vertex : graph.get(start)) {
            final LinkedList<List<String>> list = new LinkedList<>();
            dfs(vertex, graph, new HashSet<>(), list, new LinkedList<>());
            for (final List<String> str : list) {
                System.out.println(str);
            }
        }
        System.out.println(cnt);
    }

    private static void dfs(final Vertex vertex, final Map<Vertex, Set<Vertex>> graph, final Set<Vertex> cache,
            final LinkedList<List<String>> list, final LinkedList<String> current) {
        if (cache.contains(vertex)) {
            return;
        }
        if ("end".equals(vertex.node)) {
            current.add("end");
            cnt++;
            list.add(new ArrayList<>(current));
            current.removeLast();
            return;
        }
        if (vertex.isLower) {
            cache.add(vertex);
        }
        current.add(vertex.node + "->");
        for (final Vertex next : graph.get(vertex)) {
            dfs(next, graph, cache, list, current);
        }
        current.removeLast();
        if (vertex.isLower) {
            cache.remove(vertex);
        }
    }

    static class Vertex {
        final String node;
        final boolean isLower;

        Vertex(final String node) {
            this.node = node;
            boolean isLower = false;
            for (int i = 0; i < node.length(); i++) {
                if (Character.isLowerCase(node.charAt(i))) {
                    isLower = true;
                    break;
                }
            }
            this.isLower = isLower;
        }

        @Override
        public boolean equals(final Object o) {
            if (this == o)
                return true;
            if (o == null || this.getClass() != o.getClass())
                return false;
            final Vertex edge = (Vertex) o;
            return this.isLower == edge.isLower && Objects.equals(this.node, edge.node);
        }

        @Override
        public int hashCode() {
            return Objects.hash(this.node, this.isLower);
        }

        @Override
        public String toString() {
            return this.node;
        }
    }
}
