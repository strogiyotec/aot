
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;
import java.util.Stack;
import java.util.List;

public final class Main {

	public static void main(final String[] args) throws Exception {
		final Map<Character, Integer> operators = new HashMap<>(16);
		operators.put('+', 3);
		operators.put('-', 3);
		operators.put('*', 3);
		operators.put('(', 1);
		operators.put(')', 1);
		List<String> lines = Files.readAllLines(Paths.get("input"));
		long result = 0;
		for (String line : lines) {
			final Deque<String> postfix = Main.toPostfix(operators,
					line.replace("(", "( ").replace(")", " )"));
			result += Main.evaluatePostfix(postfix);
		}
		System.out.println("Part 1 " + result);
		operators.put('*', 2);
		result = 0;
		for (String line : lines) {
			final Deque<String> postfix = Main.toPostfix(operators,
					line.replace("(", "( ").replace(")", " )"));
			result += Main.evaluatePostfix(postfix);
		}
		System.out.println("Part 2 " + result);
	}

	private static long evaluatePostfix(final Deque<String> deque) {
		final Stack<String> storage = new Stack<>();
		for (final String item : deque) {
			if (Character.isDigit(item.charAt(0))) {
				storage.push(item);
			} else {
				switch (item.charAt(0)) {
					case '+':
						final long sum = Long.parseLong(storage.pop())
								+ Long.parseLong(storage.pop());
						storage.push(String.valueOf(sum));
						break;
					case '-':
						final long diff = Long.parseLong(storage.pop())
								- Long.parseLong(storage.pop());
						storage.push(String.valueOf(diff));
						break;
					case '*':
						final long mult = Long.parseLong(storage.pop())
								* Long.parseLong(storage.pop());
						storage.push(String.valueOf(mult));
						break;
					default:
						throw new IllegalStateException("Unexpected value: "
								+ item.charAt(0));
				}
			}
		}
		return Long.parseLong(storage.pop());
	}

	private static Deque<String> toPostfix(final Map<Character, Integer> operators,
			final String line) {
		final Stack<Character> stack = new Stack<>();
		final Deque<String> postfix = new ArrayDeque<>(16);
		for (final String expression : line.split("\\s")) {
			if (!expression.trim().isEmpty()) {
				if (Character.isDigit(expression.charAt(0))) {
					postfix.addLast(expression);
				} else if (expression.charAt(0) == '(') {
					stack.push('(');
				} else if (expression.charAt(0) == ')') {
					while (stack.peek() != '(') {
						postfix.addLast(stack.pop().toString());
					}
					stack.pop();
				} else if (operators.containsKey(expression.charAt(0))) {
					if (!stack.empty()) {
						final int value =
								operators.get(expression.charAt(0));
						while (!stack.isEmpty() && operators
								.get(stack.peek()) >= value) {
							postfix.addLast(stack.pop().toString());
						}
					}
					stack.push(expression.charAt(0));
				}
			}
		}
		while (!stack.isEmpty()) {
			postfix.addLast(stack.pop().toString());
		}
		return postfix;
	}
}

