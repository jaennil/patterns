import java.util.ArrayList;

class PizzeriaMenuIterator implements MenuIterator {
	private ArrayList<String> menu;
	private int position = 0;

	public PizzeriaMenuIterator(ArrayList<String> menu) {
		this.menu = menu;
	}

	@Override
	public boolean hasNext() {
		return position < menu.size();
	}

	@Override
	public String next() {
		return menu.get(position++);
	}

	@Override
	public void removeLast() {
		if (!menu.isEmpty()) {
			menu.remove(menu.size() - 1);
		}
	}

	@Override
	public void reset() {
		position = 0;
	}

}
