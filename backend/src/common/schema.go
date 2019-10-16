package common

// Database table and column names for enumeration and misc use.
var (
	DatabaseTableNames = []string{
		"product",
		"category",
		"server",
		"order",
		"table",
		"cust_code",
		"order_item",
		"settings",
	}

	ProductTableCols = []string{
		"id",
		"name",
		"desc",
		"picture",
		"price",
		"category",
		"ws_cost",
		"num_of_sides",
	}

	CategoryTableCols = []string{
		"id",
		"name",
	}

	ServerTableCols = []string{
		"id",
		"name",
		"code",
	}

	OrderTableCols = []string{
		"id",
		"starttime",
		"endtime",
		"table",
		"server",
	}

	TableTableCols = []string{
		"id",
		"table_num",
	}

	CustCodeTableCols = []string{
		"id",
		"starttime",
		"endtime",
		"code",
		"order",
	}

	OrderItemTableCols = []string{
		"id",
		"products",
		"order",
	}

	SettingTableCols = []string{
		"id",
		"value",
	}
)
