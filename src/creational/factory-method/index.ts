import { TransportFactory, SeaTransportFactory, RoadTransportFactory } from './creator/creator';

const orderAndDelivery = (factory: TransportFactory, orderList: string[]) => {
    orderList.forEach((target) => {
        const transport = factory.registerTransport();
        transport.order();
        transport.delivery(target);
    })
}

const seaTransportFactory = new SeaTransportFactory();
const roadTransportFactory = new RoadTransportFactory();

orderAndDelivery(seaTransportFactory, ["desk", "tv", "chair"]);
console.log(seaTransportFactory.checkTransportCount());

orderAndDelivery(roadTransportFactory, ["bookshelf", "refrigerator"]);
console.log(roadTransportFactory.checkTransportCount());
