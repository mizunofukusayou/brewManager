import { useEffect, useState } from "react";
interface Package {
    id: number;
    category_name: string;
    name: string;
    notes: string;
}

export default function ListPackages() {
    const [packages, setPackages] = useState<Package[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchPackages = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await fetch(
                    "http://localhost:5173/api/getpackages",
                );
                if (!response.ok) {
                    throw new Error("データ取得に失敗しました");
                }
                const data: { packages: Package[] } = await response.json();
                setPackages(data.packages);
            } catch (err) {
                const message = err instanceof Error ? err.message : "予期せぬエラーが発生しました";
                setError(message);
            } finally {
                setLoading(false);
            }
        };

        fetchPackages();
    }, []);

    if (loading) {
        return <div>ロード中...</div>;
    }

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <table style={{ width: "80%", margin: "0 auto", borderCollapse: "collapse", borderBottom: "1px solid #ddd" }}>
            <thead>
                <tr style={{borderBottom: "2px solid #ddd"}}>
                    <th scope="col" style={{ width: "30%" }}>Category</th>
                    <th scope="col" style={{ width: "40%" }}>Name</th>
                    <th scope="col" style={{ width: "30%" }}>Notes</th>
                </tr>
            </thead>
            <tbody>
                {packages.map((pkg) => (
                    <tr key={pkg.id} style={{ borderBottom: "1px solid #ddd" }}>
                        <td>{pkg.category_name}</td>
                        <td>{pkg.name}</td>
                        <td>{pkg.notes}</td>
                    </tr>
                ))}
            </tbody>
        </table>
        
    );
}

