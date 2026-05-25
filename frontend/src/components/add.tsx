import { useState } from "react";

export default function Add() {
    const l = ["package", "category"];
    const [index, setIndex] = useState(0);

    const change = () => {
        setIndex((prev) => (prev + 1) % l.length);
    };

    return (
        <div>
            <button onClick={change}>{l[(index+1) % l.length]}の追加に切り替える</button>
            {l[index] === "package" ? (
                <PackageForm />
            ) : (
                <CategoryForm />
            )}
        </div>
    );
}

function PackageForm() {
    const [selectedCategory, setSelectedCategory] = useState("");
    const categories = [
        { id: "1", name: "Category 1" },
        { id: "2", name: "Category 2" },
        { id: "3", name: "Category 3" },
    ];

    return (
        <div>
            <h2>Packageの追加</h2>
            <form action="/api/add" method="POST">
                <input type="text" placeholder="名前" name="name" required />
                <select value={selectedCategory} onChange={(e) => setSelectedCategory(e.target.value)} name="category" required>
                    <option>
                        カテゴリーを選択
                    </option>
                    {categories.map((category) => (
                        <option key={category.id} value={category.id}>
                            {category.name}
                        </option>
                    ))}
                </select>
                <button type="submit">追加</button>
                {/* Task:バックエンド側へのデータの送信の仕組みを実装(usestateで値を読み取り、手動で送信が良さそう？) */}
            </form>
        </div>
    );
}

function CategoryForm() {
    return (
        <div>
            <h2>Categoryの追加</h2>
            {/* Categoryの追加フォームをここに実装 */}
        </div>
    );
}