
import java.io.IOException;
import java.math.BigInteger;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashMap;
import java.util.Map;

public final class Day14 {

    public static void main(String[] args) throws IOException {
        final String[] lines = Files.readString(Path.of("input.txt")).split("\n");
        var rules = rules(lines);
        var template = lines[0];
        Map<String, BigInteger> freq = new HashMap<>();
        for (int i = 0; i < template.length() - 1; i++) {
            freq.put(template.substring(i, i + 2), BigInteger.ONE);
        }
        Map<Character, BigInteger> charCnt = null;
        for (int i = 0; i < 40; i++) {
            final Map<String, BigInteger> nextFreq = new HashMap<>();
            charCnt = new HashMap<>();
            for (final var entry : freq.entrySet()) {
                final String key = entry.getKey();
                final BigInteger cnt = entry.getValue();
                final String left = String.format("%c%c", key.charAt(0), rules.get(key));
                final String right = String.format("%c%c", rules.get(key), key.charAt(1));
                nextFreq.merge(left, cnt, BigInteger::add);
                nextFreq.merge(right, cnt, BigInteger::add);
                charCnt.merge(key.charAt(0), cnt, BigInteger::add);
                charCnt.merge(rules.get(key), cnt, BigInteger::add);
            }
            freq = nextFreq;
        }
        charCnt.computeIfPresent(template.charAt(template.length() - 1), (key, value) -> value.add(BigInteger.ONE));
        System.out.println(
            charCnt.values().stream().max(BigInteger::compareTo).get().subtract(charCnt.values().stream().min(BigInteger::compareTo).get())
        );
    }

    private static Map<String, Character> rules(final String[] lines) {
        final Map<String, Character> rules = new HashMap<>(lines.length << 1);
        for (int i = 2; i < lines.length; i++) {
            final String line = lines[i];
            final String[] parts = line.split(" -> ");
            rules.put(parts[0], parts[1].charAt(0));
        }
        return rules;
    }
}

