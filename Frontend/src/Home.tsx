import Navbar from "./Navbar";

function Home(): JSX.Element {
    return (
        <>
            <Navbar />
            {/* Here goes the slider juju */}

            <div>
                <h1 className="text-3xl font-extrabold font-sans text-white m-8">Videos</h1>
            </div>

        </>
    )
}

export default Home;