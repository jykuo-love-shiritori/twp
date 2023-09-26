import BoughtPage from '../../../components/BoughtPage'
import goodsData from '../../../pages/goods/goodsData.json';
import historyData from '../../../pages/cart/boughtData.json'
import NotFound from '../../../components/NotFound';

const HistoryEach = () => {
    // get the id from router
    const id = (window.location.href).slice(-1);

    // find the record if the id matches
    const record = historyData.find((item) => item.recordID.toString() === id);
    const title = record?.time + " : " + (record?.items.length)?.toString() + " items";
    let Total = 0;

    if (record != undefined) {
        Total = record.items.reduce((accumulator, item) => {
            const foundItem = goodsData.find((goods) => goods.id === item.item_id);

            if (foundItem) {
                const subtotal = item.quantity * foundItem.price;
                accumulator += subtotal;
            }
            return accumulator;
        }, 0);

        return (
            <div>
                <BoughtPage title={title} cartContainer={record?.items} total={Total} isCart={false} isButtonNeeded={false} />
            </div>
        )
    }
    else {
        return (
            <NotFound />
        )
    }
}

export default HistoryEach
