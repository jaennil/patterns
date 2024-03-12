public class MenuItem implements IMenuItem {

	private String name;

	public MenuItem(String name) {
		this.name = name;
	}

	@Override
	public String getName() {
		return name;
	}

	@Override
	public void print() {
		System.out.println(name);
	};
}
