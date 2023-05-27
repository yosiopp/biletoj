function Header() {
  return (
    <header>
      <div className="header pb-1 border-b">
        <h1 className="text-2xl px-2 inline mr-4">biletojy</h1>

        <nav className="inline">
          <ul className="inline-flex">
            <li>
              <a className="mx-2" href="#">
                tickets
              </a>
            </li>
            <li>
              <a className="mx-2" href="#">
                tags
              </a>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  );
}

export default Header;
