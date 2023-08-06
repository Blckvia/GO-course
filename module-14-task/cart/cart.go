package cart

import (
	"errors"
	"os/user"
	"test/product"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

type Cart struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
    user.User
    Items        []Item
    CurrencyCode string
    isLocked     bool
}

type Item struct {
    product.Product
    Quantity uint8
}

func TestTotalPrice(t *testing.T) {
    items := []Item{
        {
            Product: product.Product{
                ID:    "p-1254",
                Name:  "Product test",
                Price: money.New(1000, "EUR"),
            },
            Quantity: 2,
        },
        {
            Product: product.Product{
                ID:    "p-1255",
                Name:  "Product test 2",
                Price: money.New(2000, "EUR"),
            },
            Quantity: 1,
        },
    }

    c := Cart{
        ID:        "123",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        User: user.User{},
        Items: items,
        CurrencyCode: "EUR",
    }

    actual, err := c.TotalPrice()
    assert.NoError(t, err)
    assert.Equal(t, money.New(4000, "EUR"), actual)
}

func TestLock(t *testing.T) {
    c := Cart{
        ID: "123",
    }
    err := c.Lock()
    assert.NoError(t, err)
    assert.True(t, c.isLocked)
    assert.True(t, c.lockedAt.Unix() > 0)
}

func TestAlreadyLock(t *testing.T) {
    c := Cart{
        ID: "123",
        isLocked: true,
    }
    err := c.Lock()
    assert.Error(t, err)
}

func (c *Cart) TotalPrice() (*money.Money, error) {
    total := money.New(0, c.CurrencyCode)
    // var err error
    for _, v := range c.Items {
        itemSubtotal := v.Product.Price.Multiply(int64(v.Quantity))
        total, err := total.Add(itemSubtotal)
        if err != nil {
            return nil, err
        }
        
    return total, nil
    }

    return nil, nil
}

func (c *Cart) Lock() error {
    if c.isLocked {
        return errors.New("cart is already locked")
    }
    c.isLocked = true
    c.lockedAt = time.Now()
    return nil
}

func (c *Cart) delete() error {
    // to implement
    return nil
}