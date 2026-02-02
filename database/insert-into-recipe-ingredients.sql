-- SQLite
-- 1. BREAKFAST (Завтраки)
-- Овсянка с бананом
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Овсяные хлопья'), 80 FROM recipes WHERE name_ru='Овсянка с бананом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Молоко 2.5%'), 150 FROM recipes WHERE name_ru='Овсянка с бананом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Банан'), 80 FROM recipes WHERE name_ru='Овсянка с бананом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 15 FROM recipes WHERE name_ru='Овсянка с бананом';

-- Классический омлет
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яйцо куриное'), 150 FROM recipes WHERE name_ru='Классический омлет';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Молоко 2.5%'), 50 FROM recipes WHERE name_ru='Классический омлет';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Хлеб цельнозерновой'), 50 FROM recipes WHERE name_ru='Классический омлет';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Классический омлет';

-- Гречка с кефиром
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Гречневая крупа'), 80 FROM recipes WHERE name_ru='Гречка с кефиром';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Кефир 2.5%'), 250 FROM recipes WHERE name_ru='Гречка с кефиром';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 20 FROM recipes WHERE name_ru='Гречка с кефиром';

-- Творог с медом и орехами
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Творог 5%'), 150 FROM recipes WHERE name_ru='Творог с медом и орехами';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 30 FROM recipes WHERE name_ru='Творог с медом и орехами';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яблоко'), 100 FROM recipes WHERE name_ru='Творог с медом и орехами';

-- Авокадо-тост с яйцом
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Хлеб цельнозерновой'), 60 FROM recipes WHERE name_ru='Авокадо-тост с яйцом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Авокадо'), 60 FROM recipes WHERE name_ru='Авокадо-тост с яйцом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яйцо куриное'), 50 FROM recipes WHERE name_ru='Авокадо-тост с яйцом';

-- Сырники простые
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Творог 5%'), 150 FROM recipes WHERE name_ru='Сырники простые';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яйцо куриное'), 50 FROM recipes WHERE name_ru='Сырники простые';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Овсяные хлопья'), 40 FROM recipes WHERE name_ru='Сырники простые';

-- 2. SECOND_BREAKFAST (Перекусы)
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Хлеб цельнозерновой'), 60 FROM recipes WHERE name_ru='Бутерброд с яйцом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яйцо куриное'), 50 FROM recipes WHERE name_ru='Бутерброд с яйцом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Огурец'), 50 FROM recipes WHERE name_ru='Бутерброд с яйцом';

INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Овсяные хлопья'), 60 FROM recipes WHERE name_ru='Овсяное печенье';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Банан'), 60 FROM recipes WHERE name_ru='Овсяное печенье';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 15 FROM recipes WHERE name_ru='Овсяное печенье';

INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Творог 5%'), 150 FROM recipes WHERE name_ru='Творожный перекус';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яблоко'), 100 FROM recipes WHERE name_ru='Творожный перекус';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 10 FROM recipes WHERE name_ru='Творожный перекус';

INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Кефир 2.5%'), 250 FROM recipes WHERE name_ru='Йогурт с бананом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Банан'), 100 FROM recipes WHERE name_ru='Йогурт с бананом';

INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Хлеб цельнозерновой'), 60 FROM recipes WHERE name_ru='Бутерброд с сыром';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Сыр Чеддер'), 30 FROM recipes WHERE name_ru='Бутерброд с сыром';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Огурец'), 50 FROM recipes WHERE name_ru='Бутерброд с сыром';

-- 3. LUNCH (Обеды)
-- Курица с гречкой
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Куриное филе'), 150 FROM recipes WHERE name_ru='Курица с гречкой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Гречневая крупа'), 80 FROM recipes WHERE name_ru='Курица с гречкой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Курица с гречкой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Морковь'), 50 FROM recipes WHERE name_ru='Курица с гречкой';

-- Рис с овощами (Добавил курицу и масло!)
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Рис белый'), 80 FROM recipes WHERE name_ru='Рис с овощами';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Куриное филе'), 100 FROM recipes WHERE name_ru='Рис с овощами';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Морковь'), 50 FROM recipes WHERE name_ru='Рис с овощами';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Рис с овощами';

-- Чечевичный суп (Добавил говядину!)
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Чечевица'), 70 FROM recipes WHERE name_ru='Чечевичный суп';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Говядина (вырезка)'), 100 FROM recipes WHERE name_ru='Чечевичный суп';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Морковь'), 50 FROM recipes WHERE name_ru='Чечевичный суп';

-- Макароны с грудкой
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Макароны'), 100 FROM recipes WHERE name_ru='Макароны с грудкой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Куриное филе'), 120 FROM recipes WHERE name_ru='Макароны с грудкой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Макароны с грудкой';

-- Говядина с картофелем
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Говядина (вырезка)'), 150 FROM recipes WHERE name_ru='Говядина с картофелем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Картофель'), 150 FROM recipes WHERE name_ru='Говядина с картофелем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Морковь'), 50 FROM recipes WHERE name_ru='Говядина с картофелем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Говядина с картофелем';

-- Гороховая каша с мясом
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Горох'), 100 FROM recipes WHERE name_ru='Гороховая каша с мясом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Говядина (вырезка)'), 120 FROM recipes WHERE name_ru='Гороховая каша с мясом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Гороховая каша с мясом';

-- 4. SNACK (Полдники)
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Кефир 2.5%'), 250 FROM recipes WHERE name_ru='Кефир с яблоком';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яблоко'), 100 FROM recipes WHERE name_ru='Кефир с яблоком';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 15 FROM recipes WHERE name_ru='Кефир с яблоком';

INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Творог 5%'), 150 FROM recipes WHERE name_ru='Творог с огурцом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Огурец'), 100 FROM recipes WHERE name_ru='Творог с огурцом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Творог с огурцом';

-- 5. DINNER (Ужины)
-- Минтай с картофелем
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Минтай (рыба)'), 200 FROM recipes WHERE name_ru='Минтай с картофелем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Картофель'), 150 FROM recipes WHERE name_ru='Минтай с картофелем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 15 FROM recipes WHERE name_ru='Минтай с картофелем';

-- Грудка с салатом
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Куриное филе'), 150 FROM recipes WHERE name_ru='Грудка с салатом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Капуста'), 150 FROM recipes WHERE name_ru='Грудка с салатом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 15 FROM recipes WHERE name_ru='Грудка с салатом';

-- Печень куриная с рисом
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Печень куриная'), 150 FROM recipes WHERE name_ru='Печень куриная с рисом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Рис белый'), 60 FROM recipes WHERE name_ru='Печень куриная с рисом';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Печень куриная с рисом';

-- Картофельное пюре с рыбой
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Картофель'), 150 FROM recipes WHERE name_ru='Картофельное пюре с рыбой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Минтай (рыба)'), 150 FROM recipes WHERE name_ru='Картофельное пюре с рыбой';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Молоко 2.5%'), 50 FROM recipes WHERE name_ru='Картофельное пюре с рыбой';

-- Салат с креветками
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Креветки'), 150 FROM recipes WHERE name_ru='Салат с креветками';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Авокадо'), 80 FROM recipes WHERE name_ru='Салат с креветками';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Капуста'), 100 FROM recipes WHERE name_ru='Салат с креветками';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 15 FROM recipes WHERE name_ru='Салат с креветками';

-- Запеченная говядина
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Говядина (вырезка)'), 180 FROM recipes WHERE name_ru='Запеченная говядина';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Огурец'), 100 FROM recipes WHERE name_ru='Запеченная говядина';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Запеченная говядина';

-- Витаминный салат
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Капуста'), 150 FROM recipes WHERE name_ru='Витаминный салат';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Морковь'), 50 FROM recipes WHERE name_ru='Витаминный салат';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 15 FROM recipes WHERE name_ru='Витаминный салат';

-- Салат из свеклы
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Свекла'), 150 FROM recipes WHERE name_ru='Салат из свеклы';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 15 FROM recipes WHERE name_ru='Салат из свеклы';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 10 FROM recipes WHERE name_ru='Салат из свеклы';

-- Салат из моркови и яблок
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Морковь'), 100 FROM recipes WHERE name_ru='Салат из моркови и яблок';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яблоко'), 100 FROM recipes WHERE name_ru='Салат из моркови и яблок';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Оливковое масло'), 10 FROM recipes WHERE name_ru='Салат из моркови и яблок';

-- Кефир с овсяным печеньем (используем компоненты печенья)
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Кефир 2.5%'), 200 FROM recipes WHERE name_ru='Кефир с овсяным печеньем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Овсяные хлопья'), 40 FROM recipes WHERE name_ru='Кефир с овсяным печеньем';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Банан'), 40 FROM recipes WHERE name_ru='Кефир с овсяным печеньем';

-- Горсть орехов и яблоко
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Грецкий орех'), 30 FROM recipes WHERE name_ru='Горсть орехов и яблоко';
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, weight_g) SELECT id, (SELECT id FROM ingredients WHERE name_ru='Яблоко'), 150 FROM recipes WHERE name_ru='Горсть орехов и яблоко';