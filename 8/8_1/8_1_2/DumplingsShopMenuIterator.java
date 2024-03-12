import java.util.Hashtable;

class DumplingsShopMenuIterator implements MenuIterator {
	private Hashtable<Integer, String> menu;
	private int position = 0;

	public DumplingsShopMenuIterator(Hashtable<Integer, String> menu) {
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
