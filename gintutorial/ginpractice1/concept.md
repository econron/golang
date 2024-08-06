## 何を想定して作るか？

- 京都のお店を外国人に宣伝するウェブサイトが良いなと。
- お店の運営の方たちは自分で広告内容を打ち込めます。
- 外国人は英語でそれを眺められます。
- 彼らはお気に入りの店をチェックでき、チェックしたお店を閲覧可能です。
- また、そのお店の位置情報で最短ルートを構築もできます。
- 場合によってはお店の予約も自動で取れます。

## gpt4-oに仮置きで作ってもらったドメイン

### Entity

### Shop

- id (ShopId)
- name (Name)
- postal_code (PostalCode)
- city (City)
- street (Street)
- building (Building)
- phone_number (PhoneNumber)
- business_hours (BusinessHours)
- description (Description)
- latitude (Latitude)
- longitude (Longitude)
- advertisement_content (AdvertisementContent)
- favorites_count (FavoritesCount)

### User

- id (UserId)
- name (Name)
- email (Email)
- favorite_shops (FavoriteShops)

### Reservation

- id (ReservationId)
- shop_id (ShopId)
- user_id (UserId)
- reservation_datetime (ReservationDateTime)
- status (Status)

### バリューオブジェクト

### Coordinates

- latitude (Latitude)
- longitude (Longitude)

### Address

- postal_code (PostalCode)
- city (City)
- street (Street)
- building (Building)

### リポジトリ

### ShopRepository

- findById(shopId: ShopId): Shop
- findAll(): List<Shop>
- save(shop: Shop): void
- delete(shopId: ShopId): void

### UserRepository

- findById(userId: UserId): User
- findAll(): List<User>
- save(user: User): void
- delete(userId: UserId): void

### ReservationRepository

- findById(reservationId: ReservationId): Reservation
- findByShopId(shopId: ShopId): List<Reservation>
- findByUserId(userId: UserId): List<Reservation>
- save(reservation: Reservation): void
- delete(reservationId: ReservationId): void

### ユースケース
### AddShop
execute(command: AddShopCommand): void

### UpdateShop
execute(command: UpdateShopCommand): void

### DeleteShop
execute(command: DeleteShopCommand): void

### AddFavoriteShop
execute(command: AddFavoriteShopCommand): void

### RemoveFavoriteShop
execute(command: RemoveFavoriteShopCommand): void

### MakeReservation
execute(command: MakeReservationCommand): void

### コマンド
### AddShopCommand

- name (Name)
- postal_code (PostalCode)
- city (City)
- street (Street)
- building (Building)
- phone_number (PhoneNumber)
- business_hours (BusinessHours)
- description (Description)
- latitude (Latitude)
- longitude (Longitude)
- advertisement_content (AdvertisementContent)

### UpdateShopCommand

- shop_id (ShopId)
- name (Name)
- postal_code (PostalCode)
- city (City)
- street (Street)
- building (Building)
- phone_number (PhoneNumber)
- business_hours (BusinessHours)
- description (Description)
- latitude (Latitude)
- longitude (Longitude)
- advertisement_content (AdvertisementContent)

### DeleteShopCommand

- shop_id (ShopId)

### AddFavoriteShopCommand
- user_id (UserId)
- shop_id (ShopId)

### RemoveFavoriteShopCommand

- user_id (UserId)
- shop_id (ShopId)

### MakeReservationCommand

- shop_id (ShopId)
- user_id (UserId)
- reservation_datetime (ReservationDateTime)

## テーブル設計

shops

```
CREATE TABLE shops (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    city VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    building VARCHAR(100),
    phone_number VARCHAR(20) NOT NULL,
    business_hours VARCHAR(255) NOT NULL,
    description TEXT,
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    advertisement_content TEXT,
    favorites_count INT DEFAULT 0
);
```

users

```
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);
```

reservations

```
CREATE TABLE reservations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    shop_id INT NOT NULL,
    user_id INT NOT NULL,
    reservation_datetime DATETIME NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (shop_id) REFERENCES shops(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

```

favorite_shops

```
CREATE TABLE favorite_shops (
    user_id INT NOT NULL,
    shop_id INT NOT NULL,
    PRIMARY KEY (user_id, shop_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (shop_id) REFERENCES shops(id)
);

```