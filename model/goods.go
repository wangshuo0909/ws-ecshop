package model

import (
	"database/sql"
	
)

type Goods struct{
	ID int  
	Name string

    Price float64//
    Brand_id *int32//
    Category_id sql.NullInt64//

    Recommended int //是否推荐',
    Sale_vonume int //商品销量',
    Created_at string//
    UpdatedAt string `db:"updated_at"`//
}